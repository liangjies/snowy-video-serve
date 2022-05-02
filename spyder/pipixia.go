package spyder

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bitly/go-simplejson"
)

const url = "http://3363389.xyz/test.json"

func Pipixia() {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	res, err := simplejson.NewJson([]byte(body))
	if err != nil {
		fmt.Println(err)
		return
	}

	datas, err := res.Get("data").Get("data").Array()
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, _ := range datas {
		data := res.Get("data").Get("data").GetIndex(i)
		fmt.Println(data.Get("item").Get("content"))  // 标题
		fmt.Println(data.Get("item").Get("duration")) //时长
		// fmt.Println(data.Get("item").Get("video").Get("video_download").Get("url_list").GetIndex(0).Get("url")) // URL
		fmt.Println(data.Get("item").Get("video").Get("video_high").Get("url_list").GetIndex(0).Get("url"))
		// return
	}

}
