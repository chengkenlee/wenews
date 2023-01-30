package main

import (
	"encoding/json"
	"fmt"
	"github.com/nosixtools/solarlunar"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
	"wenews/util"
)

func main() {
	util.Parm()
	util.Loggers()
	var n util.News
	var arr []string
	url := "http://bjb.yunwj.top/php/60miao/qq.php"
	body := Url(url)
	err := json.Unmarshal(body, &n)
	if err != nil {
		util.Logger.Error(err.Error())
		return
	}
	for _, s := range n.Wb {
		arr = append(arr, strings.Join(s, ""))
	}
	//util.Logger.Info(strings.Join(arr, "\n"))
	wechatSend(strings.Join(arr, "\n"))
}

/*GET*/
func Url(url string) []byte {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		util.Logger.Error(err.Error())
		return nil
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
	response, err := client.Do(request)
	if err != nil {
		util.Logger.Error(err.Error())
		return nil
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		util.Logger.Error(err.Error())
		return nil
	}
	return body
}

func wechatSend(msg string) {
	s := util.Config.GetStringSlice("token.send")
	r := util.Config.GetStringSlice("token.room")

	if len(s) != 0 {
		for _, h := range s {
			util.Logger.Info(fmt.Sprintf("推送到个人，token：%s", h))
			send(msg, h, "send")
		}
	}
	if len(r) != 0 {
		for _, h := range r {
			util.Logger.Info(fmt.Sprintf("推送到群组，token：%s", h))
			send(msg, h, "room")
		}
	}
	util.Logger.Info("推送完成")
}

func send(msg, token, action string) {
	m := fmt.Sprintf("【微语简报】每天一分钟，知晓天下事！\n%s 星期%d %s\n\n%s", time.Now().Format("2006年01月02日"), int(time.Now().Weekday()), solarlunar.SolarToChineseLuanr(time.Now().Format("2006-01-02")), msg)
	escapeUrl := url.QueryEscape(m)
	u := fmt.Sprintf("http://localhost:3001/%s/%s?msg=%s", action, token, escapeUrl)
	util.Logger.Info(u)
	res := Url(u)
	util.Logger.Info(string(res))
}
