package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"
)

type Request struct {
	Method  string
	URL     string
	Cookie  string
	Headers map[string]string
	Query   map[string]string
	Params  map[string]string
}

func (request *Request) send() string {
	query := url.Values{}
	url, _ := url.Parse(request.URL)
	if request.Query != nil {
		for key, val := range request.Query {
			query.Set(key, val)
		}
		url.RawQuery = query.Encode()
	}

	var params []byte
	if request.Params != nil {
		params, _ = json.Marshal(request.Params)
	}

	req, _ := http.NewRequest(request.Method, url.String(), strings.NewReader(string(params)))

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36")
	for key, val := range request.Headers {
		req.Header.Set(key, val)
	}
	if request.Cookie != "" {
		req.Header.Set("Cookie", request.Cookie)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(res)
	// fmt.Println(string(body))
	return string(body)
}

func sendFTQQ(text, desp string) {
	request := Request{
		Method: "GET",
		URL:    "https://sc.ftqq.com/" + os.Getenv("SCKEY") + ".send",
		Query:  map[string]string{"text": text, "desp": desp},
	}
	request.send()
}

func signV2ex(ch chan<- string) {
	request := Request{
		Method: "GET",
		URL:    "https://www.v2ex.com/mission/daily/",
		Cookie: os.Getenv("V2EX_COOKIE"),
	}
	response := request.send()
	reg := regexp.MustCompile(`.*(redeem\?once=\w*).*`)
	results := reg.FindStringSubmatch(response)
	if len(results) > 0 {
		request.URL = request.URL + results[1]
		response = request.send()
		if ok, _ := regexp.MatchString(`.*(每日登录奖励已领取).*`, response); ok {
			ch <- "- v2ex成功"
			return
		}
	}
	ch <- "- v2ex失败"
	return
}

func signStudyGolang(ch chan<- string) {
	request := Request{
		Method: "GET",
		URL:    "https://studygolang.com/mission/daily/redeem",
		Cookie: os.Getenv("STUDYGOLANG_COOKIE"),
	}
	response := request.send()
	if ok, _ := regexp.MatchString(`.*(已成功领取每日登录奖励).*`, response); ok {
		ch <- "- StudyGolang 成功"
		return
	}
	ch <- "- StudyGolang 失败"
	return
}

func signGlados(ch chan<- string, num string) {
	name := "GlaDOS_" + num
	type res struct {
		Code    int         `json:"code"`
		Data    bool        `json:"data,omitempty"`
		Message string      `json:"message"`
		List    interface{} `json:"list"`
	}
	request := Request{
		Method: "POST",
		URL:    "https://glados.work/api/user/checkin",
		Cookie: os.Getenv("GLADOS_COOKIE_" + num),
		Params: map[string]string{"token": "glados_network"},
		Headers: map[string]string{
			"Authority":    "glados.work",
			"Content-Type": "application/json;charset=UTF-8",
			"Origin":       "https://glados.work",
		},
	}
	response := res{}
	err := json.Unmarshal([]byte(request.send()), &response)
	if err != nil {
		ch <- "- " + name + "失败"
		return
	}

	ch <- "- " + name + " " + response.Message
}

func signLd246(ch chan<- string) {
	type res struct {
		Msg      string `json:"msg"`
		Code     int    `json:"code"`
		UserName string `json:"userName"`
		Token    string `json:"token"`
	}

	request := Request{
		Method: "POST",
		URL:    "https://ld246.com/api/v2/login",
		Params: map[string]string{"userName": os.Getenv("LD246_NAME"), "userPassword": os.Getenv("LD246_PASSWORD")},
	}

	response := res{}
	err := json.Unmarshal([]byte(request.send()), &response)
	if err != nil {
		ch <- "- ld246 登录失败"
		return
	}
	if response.Code == 0 {
		request.Method = "GET"
		request.URL = "https://ld246.com/activity/checkin"
		request.Cookie = "symphony=" + response.Token
		request.Params = map[string]string{}
		request.send()
		body := request.send()

		reg := regexp.MustCompile(`.*今日签到获得[^>]*>([0-9]*).*`)
		results := reg.FindStringSubmatch(body)
		if len(results) > 0 {
			ch <- "- ld246获取" + results[1] + "积分"
			return
		}
	}
	ch <- "- ld246 失败"
	return
}

func sign() {
	num := 4
	desp := ""
	ch := make(chan string, num)

	go signGlados(ch, "1")
	// go signGlados(ch, "2")
	// go signGlados(ch, "3")
	go signV2ex(ch)
	go signStudyGolang(ch)
	go signLd246(ch)

	for i := 0; i < num; i++ {
		desp += <-ch + "\n"
	}
	close(ch)
	sendFTQQ(time.Now().Format("2006-01-02")+"签到", desp)
}

func main() {
	sign()
	// fmt.Println(os.Getenv("GLADOS_COOKIE_1"))
	// fmt.Println()
	// fmt.Println(os.Getenv("GLADOS_COOKIE_2"))
	// fmt.Println()
	// fmt.Println(os.Getenv("GLADOS_COOKIE_3"))
	// fmt.Println()
	// fmt.Println(os.Getenv("V2EX_COOKIE"))
	// fmt.Println()
	// fmt.Println(os.Getenv("STUDYGOLANG_COOKIE"))
	// fmt.Println()
	// fmt.Println(os.Getenv("LD246_NAME"), os.Getenv("LD246_PASSWORD"))
	// fmt.Println()
	// fmt.Println(os.Getenv("SCKEY"))
}
