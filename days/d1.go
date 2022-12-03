package days

import (
	"io"
)

type Day interface {
	Input(input io.Reader)
	Run() int
}
