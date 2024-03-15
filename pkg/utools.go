package pkg

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

type FileInfo struct {
	Name string
	Path string
}

func ParseDir(fileName string) ([]*FileInfo, error) {
	var files []*FileInfo
	err := filepath.Walk(fileName,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// 判断是不是mp4文件
			if strings.Contains(info.Name(), "mp4") {
				// 结对路径
				abs, err := filepath.Abs(path)
				if err != nil {
					return err
				}
				//fmt.Println(abs)
				files = append(files, &FileInfo{
					Name: info.Name(),
					Path: abs,
				})
			}
			return nil
		})
	if err != nil {
		log.Printf("解析目录文件失败: %v\n", err)
		return nil, err
	}
	return files, nil
}
