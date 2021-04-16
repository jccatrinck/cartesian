package mysql

import (
	"context"
	"database/sql"
	"io"

	"github.com/go-sql-driver/mysql"
	"github.com/jccatrinck/cartesian/services/points"
	"github.com/jccatrinck/cartesian/services/points/model"
	"github.com/jccatrinck/cartesian/storage/mysql/statements"
	"github.com/jmoiron/sqlx"
)

// LoadPoints implements points.Storage interface
func (m *MySQL) LoadPoints(reader io.ReadSeeker) (err error) {
	if reader == nil {
		return
	}

	walker, err := points.NewWalker(reader)

	if err != nil {
		return
	}

	err = walker.Run(func(chunk []model.Point) (err error) {
		_, err = m.db.NamedExec(statements.LoadPoints, chunk)

		if driverErr, ok := err.(*mysql.MySQLError); ok {
			if driverErr.Number == errDuplicate {
				err = nil
			}
		}

		if err != nil {
			return
		}

		return
	})

	if closer, ok := reader.(io.Closer); ok {
		err = closer.Close()

		if err != nil {
			return
		}
	}

	return
}

func (m *MySQL) getPointsByDistance(ctx context.Context, c chan<- []model.RelativePoint, e chan<- error, nstmt *sqlx.NamedStmt, p model.Point, distance int) {
	points := []model.RelativePoint{}

	err := nstmt.SelectContext(ctx, &points, model.RelativePoint{
		Point:    p,
		Distance: distance,
	})

	if err != nil && err != sql.ErrNoRows {
		e <- err
		return
	}

	c <- points

	return
}

// GetPointsByDistance implements points.Storage interface
func (m *MySQL) GetPointsByDistance(p model.Point, distance int) (result []model.RelativePoint, err error) {
	nstmtPK, err := m.db.PrepareNamed(statements.PointsByDistancePK)

	if err != nil {
		return
	}

	nstmtSpatial, err := m.db.PrepareNamed(statements.PointsByDistanceSpatial)

	if err != nil {
		return
	}

	points := make(chan []model.RelativePoint, 2)
	errors := make(chan error, 2)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go m.getPointsByDistance(ctx, points, errors, nstmtPK, p, distance)
	go m.getPointsByDistance(ctx, points, errors, nstmtSpatial, p, distance)

	select {
	case result = <-points:
		break
	case err = <-errors:
		if err != nil {
			return
		}
		break
	}

	return
}
