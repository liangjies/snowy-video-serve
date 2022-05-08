package ffmpeg

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
)

var (
	wg sync.WaitGroup
)

// 命令行调用
func Cmd(commandName string, params []string) (string, error) {
	fmt.Println("命令行调用")
	cmd := exec.Command(commandName, params...)
	//fmt.Println("Cmd", cmd.Args)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		return "", err
	}
	err = cmd.Wait()
	return out.String(), err
}

// 封面
func getVideoCover(video string, cover string, fonttxt string) error {
	// defer wg.Done()
	// 抽取一张图片，然后图片中间添加文字
	// ffmpeg -ss 0.5 -i r1.mp4 -vframes 1 -s 1080x720 -f image2 r2.jpg
	fmt.Println("Make a cover: " + video)

	vname := strings.Split(video, ".")
	videoPic := strings.Join(vname[:len(vname)-1], "") + ".jpg"

	fmt.Println(videoPic)

	cmdStr1 := fmt.Sprintf("ffmpeg -ss 0.5 -i %s -vframes 1 -s 1280x720 -f image2 %s", video, cover)
	fmt.Println("制作封面-getVideoCover命令1: " + cmdStr1)
	fmt.Println(cmdStr1)
	args := strings.Fields(cmdStr1)
	msg, err := Cmd(args[0], args[1:])
	if err != nil {
		fmt.Printf("getVideoCover1 videofailed, %v, output: %v\n", err, msg)
		return err
	}
	return err
	/*
		// 添加文字
		// ffmpeg -i r2.jpg -vf "drawtext=font='Impact':text='TEXT':fontcolor=yellow@0.8:box=1:boxcolor=red@0.2:borderw=3:fontsize-75:x=(w-tw)/2:y=((h-text_h)/2)" output2.jpg

		cmdStr2 := fmt.Sprintf("ffmpeg -i %s -vf drawtext=font='Impact':text='%s':fontcolor=yellow@0.8:box=1:boxcolor=red@0.2:borderw=3:fontsize-78:x=(w-tw)/2:y=((h-text_h)/2) %s", videoPic, fonttxt, cover)

		fmt.Println("制作封面-getVideoCover命令2: " + cmdStr2)
		fmt.Println(cmdStr2)
		args = strings.Fields(cmdStr2)
		msg, err = Cmd(args[0], args[1:])
		if err != nil {
			fmt.Printf("getVideoCover1 videofailed, %v, output: %v\n", err, msg)
			return
		}
	*/
}

func RunFFMPEG() {
	// 执行
	wg.Add(1)
	go getVideoCover("../assets/1521513705039204352.mp4", "../assets/1521513705039204352.jpg", "Music")
	wg.Wait()
}
