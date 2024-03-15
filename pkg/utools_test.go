package pkg

import (
	"testing"
)

func Test_parseDir(t *testing.T) {
	file := "/Users/yinjinghao/GolandProjects/GoBase/clip"
	dir, err := ParseDir(file)
	if err != nil {
		return
	}
	for _, v := range dir {
		t.Log(v.Path)
	}
}
