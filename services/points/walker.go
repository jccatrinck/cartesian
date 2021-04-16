package points

import (
	"encoding/json"
	"io"

	"github.com/jccatrinck/cartesian/services/points/model"
)

// WalkerProcess receives chuncks of points to process
type WalkerProcess func([]model.Point) error

// Walker stream JSON points array to be processed
type Walker struct {
	reader  io.ReadSeeker
	decoder *json.Decoder
}

// NewWalker creates a new instance of Walker
func NewWalker(reader io.ReadSeeker) (pr Walker, err error) {
	pr = Walker{
		reader: reader,
	}

	_, err = pr.reader.Seek(0, 0)

	if err != nil {
		return
	}

	pr.decoder = json.NewDecoder(pr.reader)

	// Opening bracket
	_, err = pr.decoder.Token()

	if err != nil {
		return
	}

	return
}

// Run starts decoding json as stream and calling back
func (pw Walker) Run(process WalkerProcess) (err error) {
	chunks, e := pw.run()

	for {
		select {
		case chunk, more := <-chunks:
			if len(chunk) > 0 {
				err = process(chunk)

				if err != nil {
					return
				}
			}

			if !more {
				return
			}
			break
		case err = <-e:
			if err != nil {
				return
			}
			break
		}
	}
}

func (pw Walker) newChunk() []model.Point {
	return make([]model.Point, 0, 1000)
}

func (pw Walker) run() (chunks chan []model.Point, e chan error) {
	chunks = make(chan []model.Point, 50)
	chunk := pw.newChunk()

	go func() {
		defer close(chunks)

		// While the array contains values
		for pw.decoder.More() {
			p := model.Point{}

			err := pw.decoder.Decode(&p)

			if err != nil {
				e <- err
				return
			}

			chunk = append(chunk, p)

			if len(chunk) == cap(chunk) {
				chunks <- chunk
				chunk = pw.newChunk()
			}
		}

		if len(chunk) > 0 {
			chunks <- chunk
		}
	}()

	return
}
