package excel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func ParseExcel(fileName string) ([][]string, error) {
	// 精品内容筛选 0226.xlsx
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	var rows [][]string
	for _, sheet := range f.GetSheetList() {
		ows, err := f.GetRows(sheet)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		for _, row := range ows {
			rows = append(rows, row)
		}
	}
	return rows, nil
}
