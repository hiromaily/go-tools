package main

import (
	"flag"
	"fmt"
	lg "github.com/hiromaily/golibs/log"
	"github.com/hiromaily/golibs/tmpl"
	u "github.com/hiromaily/golibs/utils"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	//GoGitDir is path for github.com directory
	GoGitDir string = "/src/github.com/"
	//GitCommandCmtID is git command to get commit ID
	GitCommandCmtID string = `git log -n 1 --pretty=format:"%H"`
)

// CmdLines is parameter for template file
type CmdLines struct {
	DirName string
	CmtID   string
}

var (
	cmdLines  = []CmdLines{}
	targetDir = flag.String("target", "", "target github.com directory path")
	usage     = `Usage: %s [options...]
Options:
  -target  path of github.com directory
e.g.:
  godepen -target ${HOME}/work/go/src/github.com
`
)

//cd ${GOPATH}/src/github.com/peterh/liner
//git checkout b850cf8c6d0ee52309aad09ac610508c6c75e819

// CheckDirectory is to check target directory
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
				cmtID, _ := exec.Command("sh", "-c", fmt.Sprintf("cd %s; ", fullPath)+GitCommandCmtID).Output()
				if err != nil {
					fmt.Println("git command failed")
				}
				//fmt.Println(fullPath)
				//fmt.Printf(" -> latest commit id is %s\n", cmtId)

				//
				dirName := strings.Replace(fullPath, os.Getenv("GOPATH"), "", 1)
				cmdLines = append(cmdLines, CmdLines{DirName: dirName, CmtID: string(cmtID)})

			} else {
				//check more deep directory
				CheckDirectory(fullPath)
			}
		}
	}

}

func init() {
	lg.InitializeLog(lg.DebugStatus, lg.LogOff, 99, "[GOTOOLS GoDependency]", "/var/log/go/gotool.log")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage, os.Args[0]))
	}

	flag.Parse()

	if *targetDir == "" {
		flag.Usage()

		os.Exit(1)
		return
	}
}

func main() {
	// targeted directory
	if *targetDir == "" {
		*targetDir = os.Getenv("GOPATH") + GoGitDir
	}
	CheckDirectory(*targetDir)

	if len(cmdLines) != 0 {
		//make sh script from template
		//fmt.Printf("\n%#v\n\n", cmdLines)
		goPath := os.Getenv("GOPATH")
		tpl, err := template.ParseFiles(goPath + "/src/github.com/hiromaily/gotools/godependency/templates/base.tpl")
		u.GoPanicWhenError(err)

		result, err := tmpl.FileTempParser(tpl, cmdLines)
		u.GoPanicWhenError(err)

		//output
		//TODO:ファイルに出力
		fmt.Println(result)
	}
}
