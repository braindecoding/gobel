package gobel

import (
	"io"
	"log"
	"reflect"

	"github.com/florianl/matf"
)

func OpenMat(filepath string) (data []float64) {
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
	//r, c, _, err = element.Dimensions()
	//data := []float64{}
	slice := reflect.ValueOf(element.Content.(matf.NumPrt).RealPart)
	for i := 0; i < slice.Len(); i++ {
		value := reflect.ValueOf(slice.Index(i).Interface()).Float()
		data = append(data, value)
	}
	return

}
