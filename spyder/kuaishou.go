package spyder

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"snowy-video-serve/core"
	"snowy-video-serve/global"
	"snowy-video-serve/initialize"
	"snowy-video-serve/model"

	// "strings"
	"github.com/bitly/go-simplejson"
	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 初始化
func init() {
	core.Viper("../config.yaml")      // 初始化Viper
	global.SYS_LOG = core.Zap()       // 初始化zap日志库
	global.SYS_DB = initialize.Gorm() // gorm连接数据库
	if global.SYS_DB != nil {
		// initialize.MysqlTables(global.SYS_DB) // 初始化表
		// 程序结束前关闭数据库链接
	}
}

// const webUrl = "https://www.kuaishou.com/graphql"
const webUrl = "http://3363389.xyz/kuaishou.json"

// const webUrl = "https://ll00.cn/Request/index.html"
const filePath = "assets"

type MyJsonName struct {
	OperationName string `json:"operationName"`
	Query         string `json:"query"`
	Variables     struct {
		HotChannelID string `json:"hotChannelId"`
		Page         string `json:"page"`
		Pcursor      string `json:"pcursor"`
	} `json:"variables"`
}

func Kuaishou() {
	//声明client 参数为默认
	client := &http.Client{}
	var someOne MyJsonName
	const jsonStream = `{"operationName":"brilliantTypeDataQuery","variables":{"hotChannelId":"00","page":"brilliant","pcursor":"1"},"query":"fragment photoContent on PhotoEntity {\n  id\n  duration\n  caption\n  likeCount\n  viewCount\n  realLikeCount\n  coverUrl\n  photoUrl\n  photoH265Url\n  manifest\n  manifestH265\n  videoResource\n  coverUrls {\n    url\n    __typename\n  }\n  timestamp\n  expTag\n  animatedCoverUrl\n  distance\n  videoRatio\n  liked\n  stereoType\n  profileUserTopPhoto\n  __typename\n}\n\nfragment feedContent on Feed {\n  type\n  author {\n    id\n    name\n    headerUrl\n    following\n    headerUrls {\n      url\n      __typename\n    }\n    __typename\n  }\n  photo {\n    ...photoContent\n    __typename\n  }\n  canAddComment\n  llsid\n  status\n  currentPcursor\n  __typename\n}\n\nfragment photoResult on PhotoResult {\n  result\n  llsid\n  expTag\n  serverExpTag\n  pcursor\n  feeds {\n    ...feedContent\n    __typename\n  }\n  webPageArea\n  __typename\n}\n\nquery brilliantTypeDataQuery($pcursor: String, $hotChannelId: String, $page: String, $webPageArea: String) {\n  brilliantTypeData(pcursor: $pcursor, hotChannelId: $hotChannelId, page: $page, webPageArea: $webPageArea) {\n    ...photoResult\n    __typename\n  }\n}\n"}`
	if err := json.Unmarshal([]byte(jsonStream), &someOne); err != nil {
		fmt.Println(err)
		return
	}

	stut, _ := json.Marshal(&someOne)
	reader := bytes.NewReader(stut)
	req, _ := http.NewRequest("GET", webUrl, reader)
	// 自定义Header
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.74 Safari/537.36 Edg/99.0.1150.46")
	req.Header.Set("Cookie", "kpf=PC_WEB; kpn=KUAISHOU_VISION; clientid=3; did=web_c395d9594374446441a18de660161d34")
	req.Header.Set("Origin", "https://www.kuaishou.com")
	req.Header.Set("Referer", "https://www.kuaishou.com/search/video?searchKey=%E6%85%A2%E6%91%87")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("accept", "*/*")
	req.Header.Set("sec-ch-ua", `"Not A;Brand";v="99", "Chromium";v="99", "Microsoft Edge";v="99"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "Windows")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))
	res, err := simplejson.NewJson([]byte(body))
	if err != nil {
		fmt.Println("e", err)
		return
	}

	feeds := res.Get("data").Get("brilliantTypeData").Get("feeds")
	feedsArray, err := feeds.Array()
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, _ := range feedsArray {
		data := feeds.GetIndex(i)
		title, _ := data.Get("photo").Get("caption").String()
		photoUrl, _ := data.Get("photo").Get("photoH265Url").String()
		// fmt.Println("title", title)
		// fmt.Println("photoUrl", photoUrl)
		err, id, videoPath, md5str := DownLoad(photoUrl)
		if err != nil {
			global.SYS_LOG.Error("视频下载失败!", zap.Any("err", err))
			continue
		}

		if err = SaveSQL(id, title, videoPath, md5str); err != nil {
			global.SYS_LOG.Error("保存数据失败!", zap.Any("err", err))
		} else {
			fmt.Println("成功下载", title)
		}
	}

}

// 视频下载
func DownLoad(imgUrl string) (err error, id uint64, videoPath string, md5str string) {
	resp, err := http.Get(imgUrl)
	if err != nil {
		return err, id, videoPath, md5str
	}
	defer resp.Body.Close()

	// 创建一个文件用于保存
	id = SnowFlake()
	fileName := fmt.Sprintf("../%s/%d.mp4", filePath, id)
	out, err := os.Create(fileName)
	if err != nil {
		return err, id, videoPath, md5str
	}
	defer out.Close()

	// md5计算
	md5Handle := md5.New()                   //创建 md5 句柄
	tr := io.TeeReader(resp.Body, md5Handle) //同时执行计算md5和复制io.Copy

	// 然后将响应流和文件流对接起来
	_, err = io.Copy(out, tr)
	if err != nil {
		return err, id, videoPath, md5str
	}

	md := md5Handle.Sum(nil)       //计算 MD5 值，返回 []byte
	md5str = fmt.Sprintf("%x", md) //将 []byte 转为 string
	// 数据重复判断
	if RepeatData(md5str) {
		if err = out.Close(); err != nil {
			return err, id, videoPath, md5str
		}
		if err = os.Remove(fileName); err != nil {
			return err, id, videoPath, md5str
		}
		return errors.New("数据重复!"), id, videoPath, md5str
	}

	// 文件路径
	videoPath = fmt.Sprintf("/%s/%d.mp4", filePath, id)
	return err, id, videoPath, md5str
}

//雪花算法
func SnowFlake() uint64 {
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	id := node.Generate()
	return uint64(id.Int64())
}

// 保存数据到SQL
func SaveSQL(id uint64, desc string, video_path string, md5str string) (err error) {
	db := global.SYS_DB.Model(&model.Videos{})
	createDate := time.Now()
	err = db.Create(&model.Videos{ID: id, UserID: 1, VideoDesc: desc, VideoPath: video_path, Status: 1, CreateTime: createDate, Hash: md5str, IsLocal: true}).Error
	if err != nil {
		global.SYS_LOG.Error("保存数据到SQL失败!", zap.Any("err", err))
		return
	}
	return
}

// 数据是否重复
func RepeatData(md5str string) bool {
	var videos model.Videos
	db := global.SYS_DB.Model(&model.Videos{})
	err := db.Where("hash= ? ", md5str).First(&videos).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
