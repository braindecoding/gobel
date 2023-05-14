package gobel

import (
	"io"
	"log"
	"reflect"

	"github.com/florianl/matf"
	"gorgonia.org/tensor"
)

func OpenMat(filepath string) (t tensor.Tensor) {
	modelfile, err := matf.Open(filepath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer matf.Close(modelfile)
	element, err := matf.ReadDataElement(modelfile)
	if err != nil && err != io.EOF {
		log.Fatal(err)
		return
	}
	r, c, _, err := element.Dimensions()
	data := []uint64{}
	slice := reflect.ValueOf(element.Content.(matf.NumPrt).RealPart)
	for i := 0; i < slice.Len(); i++ {
		value := reflect.ValueOf(slice.Index(i).Interface()).Uint()
		data = append(data, value)
	}

	t = tensor.New(tensor.WithShape(r, c), tensor.WithBacking(data))
	return

}
