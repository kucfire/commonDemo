#!/bin/sh
#FIlename: init.sh
#Function: Create a init github respositories shell
#Version: 0.0.1
#Author: kucfire
#Data: 2021-01-19

## 复制本脚本文件到对应的空白项目下并运行,会搭建一个初始的项目框架
## 需要将init.sh的权限提升到777(即执行sudo chmod -R 777 init.sh或者sudo chmod u+x init.sh)

## 判断输入参数数量,数量不达标则退出脚本
if [ $# != 2 ]; then
    echo "Please input least 2 param"
    exit 1
fi
## 判断输入的参数的格式，TODO

## 创建readme.md,并写入对应的文本框架(文本框架可自定义)
if [ ! -f readme.md ]; then
    echo "readme.md isn't exist"
    touch readme.md
    echo "created readme.md successful"
else
    echo "readme.md is exist"
    echo "clear readme.md's content"
    cat /dev/null >readme.md # 清空文件内容
fi

echo "Write format content to readme.md"
## echo "将自定义框架放入这里,最好以markdown的格式进行编写,在上传至github的时候可以呈现出对应的效果" >>readme.md
## example "hello,world!":
echo "Hello.world!" >>readme.md
## example markdown format:
## 后期有些一些关于markdown的文本时会放到这里做成示例

## 构建一个.gitignore
if [ ! -f .gitignore ]; then
    echo ".gitignore isn't exist"
    touch .gitignore
    echo "created .gitignore successful"
else
    echo ".gitignore is exist" # 这里不需要清空内容，可能原本存在的.gitignore就有屏蔽的内容
    cat /dev/null >.gitignore  # 测试用的，实际使用该脚本时需要注释掉该选项
fi
## 写入常用的禁止上传github的无用文件
echo ".history/\nlogs/\ndist/\nlogs/\nnode_modules/\npackage-lock.json\n.vscode" >>.gitignore # 产生该文件的为vscode中的一款历史记录的插件，该插件会随着保存而创建多个文件，造成大量的无需上传github的文件，影响上传速率甚至浪费资源

## 初始化项目的git环境,直接git init,不做其他判断
# if [ ! -d .git ]; then
#     echo ".git isn't exist"
#     git init
# fi
git init

## 将本地与远程仓库进行关联
## 需要先判断本地是否已经存在该远程仓库的链接
if [ ! -n $1 ]; then
    echo "please input you respossitories"
fi
# git remote -v | awk '{print $2}'
git remote add $1 $2
echo "remote add successful"

## 创建项目所需的文件分类
## example about go:
# touch -p dao/demo.go controller/demo.go dto/demo.go logs bin conf/demo.go router/demo.go
# TODO 用正则表达式筛选文件夹名称并使用echo放进对应的demo.go里面去
# mkdir dao controller dto logs bin conf router
# touch dao/demo.go controller/demo.go dto/demo.go conf/demo.conf router/demo.go
# echo "package dao" >dao/demo.go
# echo "package controller" >controller/demo.go
# echo "package dto" >dto/demo.go
# echo "package router" >router/demo.go
# cp init.sh bin
## 如果有自己的demo，可以直接clone下demo到该empty respositories中
# git clone

## 上传当前目录的所有东西到远程代码仓库
git add .
git commit -m "init respositories"
git branch -m main
git push -u $1 main
