package statements

// Schema of the database
const Schema = `
DROP TABLE IF EXISTS point;

CREATE TABLE point (
	x INT(11),
	y INT(11),
	point POINT NOT NULL SRID 0, # SRID: Cartesian all-purpose abstract plane
	PRIMARY KEY (x, y),
	SPATIAL INDEX (point)
);

DROP TRIGGER IF EXISTS point_before_insert;

CREATE TRIGGER point_before_insert
BEFORE INSERT ON point
FOR EACH ROW
SET NEW.point = ST_SRID(POINT(NEW.x, NEW.y), 0);

DROP FUNCTION IF EXISTS taxicab_circle;

CREATE FUNCTION taxicab_circle(x INT(11), y INT(11), d INT(11))
RETURNS POLYGON DETERMINISTIC
RETURN ST_SRID(POLYGON(
	LINESTRING(
		POINT(x, y+d+1e-9),
		POINT(x+d+1e-9, y),
		POINT(x, y-d-1e-9),
		POINT(x-d-1e-9, y),
		POINT(x, y+d+1e-9)
	)
), 0);
`
