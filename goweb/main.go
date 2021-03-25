package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
)
var(
	port int
	targetproject string
)
func main() {


	// 要遍历的文件 夹
	dir := `../`

	// 遍历的文件夹
	// 参数：要遍历的文件夹，层级（默认：0）
	findDir(dir, 0)

	PrintArgs1(os.Args)
}

// 遍历的文件夹
func findDir(dir string, num int) {

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
			findDir(dir+`/`+fi.Name(), num+1)
		} else {
			println(`文件：`, fi.Name())
		}
	}

}

//变参函数的定义方式
func PrintArgs1(args ...interface{}) {
	argsarr := args[0].([]string)
	fmt.Println(argsarr)
	fmt.Println(len(argsarr))
	if len(argsarr) != 3 {
		printError()
		return
	}
	if !IsNum(argsarr[2]) {
		printError()
		return
	}

	for k, v := range argsarr {
		fmt.Println(k, " =", v, reflect.TypeOf(v))
	}

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
