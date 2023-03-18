#!/usr/bin/env bash

src="./mecab-naist-jdic"
data_dir=".."
dict="./${data_dir}/njd.dict"

go run main.go -dict ${src} -out ${dict}
