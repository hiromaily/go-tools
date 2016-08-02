package main

import (
	"flag"
	"fmt"
	"github.com/hiromaily/golibs/tmpl"
	u "github.com/hiromaily/golibs/utils"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

const GO_GIT_DIR string = "/src/github.com/"
const GIT_COMMAND_CMITID string = `git log -n 1 --pretty=format:"%H"`

var cmdLines = []CmdLines{}

type CmdLines struct {
	DirName string
	CmtId   string
}

func CheckDirectory(target string) {
	// read directory
	fis, err := ioutil.ReadDir(target)
	u.GoPanicWhenError(err)

	for _, fi := range fis {
		//fmt.Println(fi.Name())
		if fi.Name() == ".git" || fi.Name() == ".idea" {
			continue
		}

		fullPath := filepath.Join(target, fi.Name())

		if fi.IsDir() {
			//if .git directory is exiting, execute git command
			if u.IsExistDir(fullPath + "/.git") {
				//get latest commit id by git command
				//out, err := exec.Command("git", "log", "-n", "1", "--pretty=format:\"%H\"").Output()
				cmtId, _ := exec.Command("sh", "-c", fmt.Sprintf("cd %s; ", fullPath)+GIT_COMMAND_CMITID).Output()
				if err != nil {
					fmt.Println("git command failed")
				}
				//fmt.Println(fullPath)
				//fmt.Printf(" -> latest commit id is %s\n", cmtId)

				//
				dirName := strings.Replace(fullPath, os.Getenv("GOPATH"), "", 1)
				cmdLines = append(cmdLines, CmdLines{DirName: dirName, CmtId: string(cmtId)})

			} else {
				//check more deep directory
				CheckDirectory(fullPath)
			}
		}
	}

}

func main() {
	var targetDir = flag.String("target", "", "target github.com directory path")
	flag.Parse()

	// targeted directory
	if *targetDir == "" {
		*targetDir = os.Getenv("GOPATH") + GO_GIT_DIR
	}
	CheckDirectory(*targetDir)

	if len(cmdLines) != 0 {
		//make sh script from template
		//fmt.Printf("\n%#v\n\n", cmdLines)
		goPath := os.Getenv("GOPATH")
		tpl, err := template.ParseFiles(goPath + "/src/github.com/hiromaily/gotools/godependency/templates/base.tpl")
		u.GoPanicWhenError(err)

		result, err := tmpl.FileTempParser(tpl, cmdLines)
		fmt.Println(result)

		//TODO:ファイルに出力
	}
}
