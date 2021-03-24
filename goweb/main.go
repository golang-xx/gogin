package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"
)

func FormatPath(s string) string {
	switch runtime.GOOS {
	case "windows":
		return strings.Replace(s, "/", "\\", -1)
	case "darwin", "linux":
		return strings.Replace(s, "\\", "/", -1)
	default:
		log.Println("only support linux,windows,darwin, but os is " + runtime.GOOS)
		return s
	}
}

func copyDir(src string, dest string) {
	src = FormatPath(src)
	dest = FormatPath(dest)
	log.Println(src)
	log.Println(dest)

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("xcopy", src, dest, "/I", "/E")
	case "darwin", "linux":
		cmd = exec.Command("cp", "-R", src, dest)
	}

	outPut, e := cmd.Output()
	if e != nil {
		log.Println(e.Error())
		return
	}
	fmt.Println(string(outPut))
}

func main() {
    copyDir("../ginweb","../ginweb2")
}