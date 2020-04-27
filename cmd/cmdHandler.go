package cmd

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
	"time"
)

//chCmd 执行结果传递的channel
type chCmd struct {
	Out    bytes.Buffer
	Stderr bytes.Buffer
	Err    error
}
//RunCmd 执行cmd命令并返回输出内容 如果执行失败 error不为nil
func RunCmd(str string) (result string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) //命令最长十秒要返回
	defer cancel()
	var ch = make(chan *chCmd)
	str += " && echo S_OK || echo E_FAIL" //判断成功失败获取输出后将这两个字符去掉
	fmt.Println("执行命令:",str)
	cmd := exec.CommandContext(ctx, "cmd", "/C", str)
	if runtime.GOOS == "windows" { //隐藏黑窗口
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	}
	defer func(cmd *exec.Cmd) {                            //回收cmd
		if cmd.ProcessState != nil || cmd.Process == nil { //如果结束了就不强行杀死
			//判断cmd进程状态，已退出进程就不进行处理
		} else {
			if err := cmd.Process.Kill(); err != nil {
				fmt.Println("kill cmd error", err)
			}
		}

	}(cmd)
	//使用协程执行命令,避免阻塞
	go func(cmd *exec.Cmd, ch chan *chCmd) {
		cmdCh := new(chCmd)
		cmd.Stdout = &cmdCh.Out
		cmd.Stderr = &cmdCh.Stderr
		cmdCh.Err = cmd.Run()
		ch <- cmdCh
	}(cmd, ch)
	var logString string
End:
	for {
		select {
		case cmdCh := <-ch:
			cancel()
			if cmdCh.Err != nil {
				//检测报错是否是因为超时引起的
				if ctx.Err() != nil && ctx.Err() == context.DeadlineExceeded {
					logString = fmt.Sprintf("命令超时错误output:%s,err:%s,stderr:%s", cmdCh.Out.String(), cmdCh.Err.Error(), cmdCh.Stderr.String())

				} else {
					logString = fmt.Sprintf("错误:output:%s,err:%s,stderr:%s", cmdCh.Out.String(), cmdCh.Err.Error(), cmdCh.Stderr.String())

				}
			} else {
				logString = cmdCh.Out.String()
			}
			//fmt.Println("返回结果")
			//fmt.Printf("%#s",logString)
			//fmt.Println("返回结果")
			if strings.Index(logString, "S_OK") > -1 {
				logString = strings.ReplaceAll(logString, "S_OK", "")

				logString = strings.TrimSpace(logString)
				return logString, nil

			} else {

				logString = strings.ReplaceAll(logString, "E_FAIL", "")
				logString = strings.TrimSpace(logString)
				return "", errors.New(logString)
			}
		case <-ctx.Done():
			cancel()
			break End
		default:
			time.Sleep(300 * time.Millisecond) //防止阻塞
		}
	}
	return "", errors.New("cmd exec failed")

}
