package main

import (
	"clip/pkg"
	"clip/pkg/excel"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sourcegraph/conc/pool"
)

type clipInfo struct {
	title       string // 标题
	mediaAssets string // 媒资ID
	mediaName   string // 媒资名称
	programID   string // 节目id
	formType    string
	fileName    string // 本地文件名字
	path        string // 本地文件路径
}

func findFex(media string, info []*pkg.FileInfo) (name, path string) {
	for _, i := range info {
		n := strings.Split(i.Name, "_")
		if strings.Contains(n[0], media) {
			return i.Name, i.Path
		}
	}
	return "", ""
}

func main() {
	// 1. 解析excel文件
	parseExcel, err := excel.ParseExcel("精品内容筛选 0226.xlsx")
	if err != nil {
		log.Printf("parse excel failed, err:%s \n", err)
		return
	}
	// 2. 遍历目录文件名称和excel文件内容对接上 ParseDir
	parseDir, err := pkg.ParseDir("/Users/yinjinghao/移动盘下载")
	//parseDir, err := pkg.ParseDir(".")
	if err != nil {
		log.Printf("parseDir to failed, err:%s\n", err)
		return
	}
	var clipCh = make(chan *clipInfo, len(parseExcel))
	defer close(clipCh)
	for _, i := range parseExcel {
		//var id = i[1]
		//if i[0] == "篮球" || i[0] == "足球" {
		//	id = i[10]
		//}
		id := i[8]
		name, path := findFex(id, parseDir)
		c := &clipInfo{
			title:       i[0],
			mediaAssets: i[1],
			mediaName:   i[2],
			programID:   i[10],
			formType:    i[11],
			fileName:    name,
			path:        path,
		}
		clipCh <- c
	}
	file, err := os.OpenFile("file.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	failedFile, err := os.OpenFile("file_err.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer failedFile.Close()
	// 3. 并发调用上传接口
	// 3.1 启用线程池，设置最大线程数
	p := pool.New().WithMaxGoroutines(8)
	for {
		select {
		case c := <-clipCh:
			p.Go(func() {
				if c.path != "" {
					c.mediaName = strings.Trim(c.mediaName, "\n")
					fmt.Printf("%v|%v|%v｜%v\n", c.mediaName, c.programID, c.mediaAssets, c.path)
					//rep, err := request.UploadFile(c.path)
					//if err != nil {
					//	// 失败了加入队列重试
					//	clipCh <- c
					//	fmt.Printf("%v调用失败，加入重试队列, err:%v\n", c.mediaAssets, err)
					//	return
					//}
					//if rep.Code != "000000" {
					//	clipCh <- c
					//	fmt.Printf("%v调用失败，加入重试队列, err:%v\n", c.mediaAssets, err)
					//	return
					//}
					//cont := fmt.Sprintf("%v|%v|%v|%v|%v|%v|%v\n",
					//	c.title,
					//	c.mediaAssets,
					//	c.mediaName,
					//	c.programID,
					//	c.formType,
					//	rep.Data[0].LocalPath,
					//	rep.Data[0].Duration,
					//)
					//fmt.Printf(cont)
					//file.WriteString(cont)
				} else if c.title != "已选内容" {
					failedFile.WriteString(fmt.Sprintf("%v|%v|%v\n", c.title, c.mediaAssets, c.mediaName))
				}
			})
		default:
			if len(clipCh) == 0 {
				p.Wait()
				fmt.Println("任务队列为空")
				return
			}
		}
	}
	// 申报逻辑
	//var data = make([][]string, 0, 300)
	//file, err := os.Open("file.txt")
	//if err != nil {
	//	log.Fatalf("open file to failed err:%v\n", err)
	//}
	//reader := bufio.NewReader(file)
	//for {
	//	readString, err := reader.ReadString('\n')
	//	if err == io.EOF {
	//		break
	//	}
	//	//s := strings.Split(readString[:-1], "|")
	//	s := strings.Split(strings.TrimSpace(readString), "|")
	//	if len(s) != 7 {
	//		log.Printf("%v strings.Split to not 7\n", s[0])
	//	}
	//	data = append(data, s)
	//}
	//for _, v := range data {
	//	i, _ := strconv.ParseFloat(v[6], 64)
	//	i = i / 1000
	//	fmt.Printf(`{
	//"publishTemplateId": "9545c159-a116-4e08-bf25-09ad7feab046",
	//"videoFileInfo": {
	//   "videoName": "%s",
	//   "filePath": "%s",
	//   "duration": %f|%d
	//}}`, v[2], v[5], i, int(pkg.Decimal(i)))
	//}
}
