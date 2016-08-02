#!/bin/sh

{{range .}}
cd ${GOPATH}{{.DirName}}
git checkout {{.CmtId}}
{{end}}