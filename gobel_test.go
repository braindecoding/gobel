package gobel

import (
	"fmt"
	"testing"
)

func TestOpenMat(t *testing.T) {
	matvar := OpenMat("stim.mat")
	fmt.Println(matvar.Shape())

}
