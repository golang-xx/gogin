package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)
var(
	port int
	targetproject string
)
func main() {

	dealArgs(os.Args)
	// 要遍历的文件 夹
	copyDir("../ginweb","../"+targetproject)
	fmt.Println("Sussess!")

	replacestr("../"+targetproject,"ginweb",targetproject)
	fmt.Println(port)
	replacestr("../"+targetproject,"8090", strconv.Itoa(port))


}

func FormatPath(s string) string {
	switch runtime.GOOS {
	case "windows":
		return strings.Replace(s, "/", "\\", -1)
	case "darwin", "linux":
		return strings.Replace(s, "\\", "/", -1)
	default:
		fmt.Println("only support linux,windows,darwin, but os is " + runtime.GOOS)
		return s
	}
}
func copyDir(src string, dest string) {
	src = FormatPath(src)
	dest = FormatPath(dest)
	fmt.Println(src)
	fmt.Println(dest)

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("xcopy", src, dest, "/I", "/E")
	case "darwin", "linux":
		cmd = exec.Command("cp", "-R", src, dest)
	}

	outPut, e := cmd.Output()
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	fmt.Println(string(outPut))
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

func replacestr(rootpath string,source string,target string)  {
	helper := ReplaceHelper{
		Root:    rootpath,
		OldText: source,
		NewText: target,
	}
	err := helper.DoWrok()
	if err == nil {
		fmt.Println("done!")
	} else {
		fmt.Println("error:", err.Error())
	}
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