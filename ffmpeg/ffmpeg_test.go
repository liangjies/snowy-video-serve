package ffmpeg

import (
	"fmt"
	"testing"
)

func TestFFMPEG(t *testing.T) {
	// RunFFMPEG()
	// 获取视频列表
	videoList, err := VideoData()
	if err != nil {
		fmt.Println("获取视频列表错误")
		return
	}

	for _, video := range videoList {
		videoPath := fmt.Sprintf("../assets/video/%d.mp4", video.ID)
		cover := fmt.Sprintf("../assets/cover/%d.jpg", video.ID)
		// 生成封面
		if err := getVideoCover(videoPath, cover, "Music"); err == nil {
			video.CoverPath = fmt.Sprintf("/assets/cover/%d.jpg", video.ID)
			// 写入数据库
			if err = UpdateCover(video); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("更新成功")
			}
		}
	}
}

func TestDuration(t *testing.T) {
	// 获取视频列表
	videoList, err := VideoDataDuration()
	if err != nil {
		fmt.Println("获取视频列表错误")
		return
	}
	for _, video := range videoList {
		videoPath := fmt.Sprintf("../assets/video/%d.mp4", video.ID)
		video.VideoSeconds = float32(GetMP4FileDuration(videoPath))

		if err = UpdateDuration(video); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("更新成功")
		}
	}
}
