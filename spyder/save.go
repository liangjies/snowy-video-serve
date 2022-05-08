package spyder

import (
	"fmt"
	"snowy-video-serve/utils/upload"
)

func UploadVideo() {
	videoPath := fmt.Sprintf("../assets/video/1523205061369401344.mp4")

	var oss upload.OSS
	oss = &upload.Qiniu{}
	// oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadLocalFile(videoPath, "1523205061369401344.mp4", "video/")
	if uploadErr != nil {
		fmt.Println(uploadErr)
	}
	fmt.Println(filePath)
	fmt.Println(key)
}
