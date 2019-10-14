#!/bin/sh

{{range .}}
cd ${GOPATH}{{.DirName}}
git checkout {{.CmtID}}
{{end}}