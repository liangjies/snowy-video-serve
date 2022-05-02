package spyder

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"snowy-video-serve/global"
	"snowy-video-serve/model"
	"time"

	"snowy-video-serve/core"
	"snowy-video-serve/initialize"

	"github.com/bwmarrin/snowflake"
)

func init() {
	core.Viper("../config.yaml")      // 初始化Viper
	global.SYS_LOG = core.Zap()       // 初始化zap日志库
	global.SYS_DB = initialize.Gorm() // gorm连接数据库
	if global.SYS_DB != nil {
		// initialize.MysqlTables(global.SYS_DB) // 初始化表
		// 程序结束前关闭数据库链接
	}
}

func RunSpyder() {
	db := global.SYS_DB.Model(&model.Videos{})
	// 雪花
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileInfoList := ReadDirFiles("G:\\MyDrivers\\hotfix\\国\\新建文件夹\\")
	for _, file := range fileInfoList {
		name, ext := Ext(file.Name())
		fmt.Println(name, ext)
		video_path := "http://192.168.10.168:8000/videosPath/" + file.Name()

		// Generate a snowflake ID.
		id := node.Generate()
		fmt.Printf("ID : %d\n", uint64(id.Int64()))
		createDate := time.Now()
		err = db.Create(&model.Videos{ID: uint64(id.Int64()), UserID: 1, VideoDesc: name, VideoPath: video_path, Status: 1, CreateTime: createDate}).Error
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}
func ReadDirFiles(path string) []fs.FileInfo {
	fileInfoList, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}

	return fileInfoList
}

func Ext(path string) (name string, ext string) {
	for i := len(path) - 1; i >= 0 && !os.IsPathSeparator(path[i]); i-- {
		if path[i] == '.' {
			return path[:i], path[i:]
		}
	}
	return name, ext
}
