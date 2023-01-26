package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	ansi "github.com/k0kubun/go-ansi"
	progressbar "github.com/schollz/progressbar/v3"
)

type FileInfo struct {
	FullPath string
	Size     int64
	IsDir    bool
}

func AnalysisDir(dirName string) ([]*FileInfo, error) {
	//fmt.Printf("anaylsis dir: %s\n", dirName)
	files := make([]*FileInfo, 0)
	err := filepath.Walk(
		dirName, func(path string, info os.FileInfo, err error) error {
			//fmt.Println(path)
			if info == nil {
				fmt.Println("walk info is null")
				return nil
			}
			f := new(FileInfo)
			if info.IsDir() {
				f.IsDir = true
			}
			f.FullPath = path
			f.Size = info.Size()
			files = append(files, f)
			return nil
		})
	return files, err
}

func CopyFile(srcName, dstName string) (written int64, err error) {
	//fmt.Printf("%s => %s\n", srcName, dstName)
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func MakeTargetPath(source, target, source_path string) (string, error) {

	///
	s := strings.Split(source_path, source)
	if len(s) == 2 {
		return target + s[1], nil
	}
	return "", fmt.Errorf("string split error")
}

func AddLog(log, item string) error {
	file, err := os.OpenFile(log, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	write.WriteString(item + "\n")
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
	return nil
}

func CheckLog(log, item string) (bool, error) {

	//fmt.Printf("check item:%s\n", item)
	file, err := os.OpenFile(log, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return false, err
	}
	//及时关闭file句柄
	defer file.Close()
	//读原来文件的内容，并且显示在终端
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		//fmt.Printf("compare %s:%s\n", str, item)
		if strings.Replace(str, "\n", "", -1) == item {
			return true, nil
		}
	}
	return false, nil
}

func DeleteFile() {

}

func MakeDir(dir string) error {
	return os.MkdirAll(dir, os.ModePerm)
}

func ToDelete(fileName string) bool {
	if strings.Contains(fileName, ".txt") {
		return true
	}
	if strings.Contains(fileName, ".html") {
		return true
	}
	if strings.Contains(fileName, ".url") {

		return true
	}
	if strings.Contains(fileName, ".torrent") {

		return true
	}
	if strings.Contains(fileName, ".js") {

		return true
	}

	if strings.Contains(fileName, "快来安装") {
		return true
	}
	if strings.Contains(fileName, "福利") {
		return true
	}
	if strings.Contains(fileName, "点击观看") {
		return true
	}
	if strings.Contains(fileName, "精彩片头") {
		return true
	}
	if strings.Contains(fileName, "prpxv.mp4") {
		return true
	}
	if strings.Contains(fileName, "gif") {
		return true
	}
	if strings.Contains(fileName, "宣传图") {
		return true
	}
	if strings.Contains(fileName, "png") {
		return true
	}
	if strings.Contains(fileName, "老司机") {
		return true
	}
	if strings.Contains(fileName, "找的到") {
		return true
	}
	if strings.Contains(fileName, "星际末世") {
		return true
	}
	if strings.Contains(fileName, "AV大平台") {
		return true
	}
	if strings.Contains(fileName, "快感上腺") {
		return true
	}

	//
	if strings.Contains(fileName, "超優質愛情動作片") {
		return true
	}
	if strings.Contains(fileName, "世界杯最方便的視頻") {
		return true
	}
	if strings.Contains(fileName, "全部免費") {
		return true
	}
	if strings.Contains(fileName, "张信哲的成功致富方法") {
		return true
	}

	return false
}

func main() {

	var source_dir string
	var log_file string
	var check bool

	flag.StringVar(&source_dir, "source", "", "source dir path,default null")
	flag.StringVar(&log_file, "log", "./log", "log file path,default ./log")
	flag.BoolVar(&check, "check", true, "check to delete, true or false")
	flag.Parse()

	if source_dir == "" {
		log.Fatal("params error")
	}

	// get source dir structure
	sourceFiles, err := AnalysisDir(source_dir)
	if err != nil {
		log.Fatalf("analysis dir error : %s\n", err.Error())
	}

	// get not dir count
	fileCount := 0
	for _, f := range sourceFiles {
		if !f.IsDir {
			fileCount++
		}
	}

	// init bar
	bar := progressbar.NewOptions(fileCount,
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(35),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[red]=[reset]",
			SaucerHead:    "[red]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))

	// copy to target dir
	for _, file := range sourceFiles {
		bar.Add(1)
		if file.FullPath == source_dir {
			//fmt.Printf("skip %s\n", source_dir)
			continue
		}

		if !ToDelete(file.FullPath) {
			continue
		}

		if !check {
			fmt.Printf("delete file %s\n", file.FullPath)
			e := os.Remove(file.FullPath)
			if e != nil {
				log.Fatalf("delete file error :%s\n", err.Error())
			}
		} else {

			fmt.Printf("check delete file %s\n", file.FullPath)
		}

		// add log
		err = AddLog(log_file, file.FullPath)
		if err != nil {
			log.Fatalf("add log error:%s\n", err.Error())
		}
	}
}
