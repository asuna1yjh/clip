package excel

import (
	"fmt"
	"testing"
)

func TestParseExcel(t *testing.T) {
	fileName := "/Users/yinjinghao/GolandProjects/GoBase/clip/精品内容筛选 0226.xlsx"
	excel, err := ParseExcel(fileName)
	if err != nil {
		t.Error(err)
	}
	for _, v := range excel {
		fmt.Printf("%v\n", v)
	}
}
