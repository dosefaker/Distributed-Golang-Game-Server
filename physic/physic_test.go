package physic

import (
	"github.com/gazed/vu/math/lin"
	"testing"
)

func TestWorld(t *testing.T) {
	w := &World{}
	w.Init()
	w.CreateEnitity("Tank", int64(123))
}
