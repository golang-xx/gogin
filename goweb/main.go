package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)
var(
	port int
	targetproject string
)
func main() {

	var cmd *exec.Cmd

	dealArgs(os.Args)
	// 要遍历的文件 夹
	dir := `../`
	cmd = exec.Command("mkdir", dir+targetproject)
	outPut, e := cmd.Output()
	if e != nil {
		fmt.Println(e.Error())
	}
	fmt.Println(string(outPut))
	// 遍历的文件夹
	// 参数：要遍历的文件夹，层级（默认：0）
	findDir(dir+"ginweb", 0,targetproject)
	fmt.Println("Sussess!")

}

// 遍历的文件夹
func findDir(dir string, num int,tarproject string) {

	 fileinfo, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	// 遍历这个文件夹
	for _, fi := range fileinfo {

		// 重复输出制表符，模拟层级结构
		print(strings.Repeat("\t", num))

		// 判断是不是目录
		if fi.IsDir() {
			println(`目录：`, fi.Name())
			findDir(dir+`/`+fi.Name(), num+1,tarproject)
		} else {
			println(`文件：`, fi.Name())
		}
	}

}

//变参函数的定义方式
func dealArgs(args ...interface{}) {
	argsarr := args[0].([]string)
	if len(argsarr) != 3 {
		printError()
		return
	}
	if !IsNum(argsarr[2]) {
		printError()
		return
	}
	targetproject = argsarr[1]
	port,_= strconv.Atoi(argsarr[2] )

}

func printError() {
	fmt.Println("Error !!!!!!!!!")
	fmt.Println("goweb  newproject newport \n eg: goweb ginweb xxx 8887 、\n")
	return
}

func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

type ReplaceHelper struct {
	Root    string //根目录
	OldText string //需要替换的文本
	NewText string //新的文本
}

func (h *ReplaceHelper) DoWrok() error {

	return filepath.Walk(h.Root, h.walkCallback)

}

func (h ReplaceHelper) walkCallback(path string, f os.FileInfo, err error) error {

	if err != nil {
		return err
	}
	if f == nil {
		return nil
	}
	if f.IsDir() {
		//fmt.Pringln("DIR:",path)
		return nil
	}

	//文件类型需要进行过滤

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		//err
		return err
	}
	content := string(buf)

	//替换
	newContent := strings.Replace(content, h.OldText, h.NewText, -1)

	//重新写入
	ioutil.WriteFile(path, []byte(newContent), 0)

	return err
}