package main

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)

func main() {
	fileName := "./my.log"
	config := tail.Config{
		ReOpen:    true,                                 //重新打开
		Follow:    true,                                 //是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 文件打开的位置
		MustExist: false,                                // 是不是必须存在
		Poll:      true,                                 //
	}
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Printf("tail file failed error:%s\n", err.Error())
		return
	}
	var (
		msg *tail.Line
		ok  bool
	)
	for {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, sfilename;%s\n", fileName)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("msg:", msg.Text)
	}
}
