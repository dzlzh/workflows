package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

func sendRequest(method, url, cookie string) string {
	req, _ := http.NewRequest(method, url, nil)

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36")
	req.Header.Set("Cookie", cookie)

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

func sendFTQQ(message string) {
	sendRequest("GET", "https://sc.ftqq.com/"+os.Getenv("SCKEY")+".send?text="+message, "")
}

func signV2ex() bool {
	baseUrl := "https://www.v2ex.com/mission/daily/"
	cookie := os.Getenv("V2EX_COOKIE")
	body := sendRequest("GET", baseUrl, cookie)
	reg := regexp.MustCompile(`.*(redeem\?once=\w*).*`)
	results := reg.FindStringSubmatch(body)
	if len(results) > 0 {
		signUrl := baseUrl + results[1]
		sendRequest("GET", signUrl, cookie)
		body = sendRequest("GET", baseUrl, cookie)

		// reg = regexp.MustCompile(`.*(每日登录奖励已领取).*`)
		// results = reg.FindStringSubmatch(body)
		if ok, _ := regexp.MatchString(`.*(每日登录奖励已领取).*`, body); ok {
			return true
		}
	}
	return false
}

func signStudyGolang() bool {
	baseUrl := "https://studygolang.com/mission/daily/redeem"
	cookie := os.Getenv("STUDYGOLANG_COOKIE")
	body := sendRequest("GET", baseUrl, cookie)
	// reg := regexp.MustCompile(`.*(已成功领取每日登录奖励).*`)
	// results := reg.FindStringSubmatch(body)
	if ok, _ := regexp.MatchString(`.*(已成功领取每日登录奖励).*`, body); ok {
		return true
	}
	return false
}

func signGlodos() bool {
	baseUrl := "https://glados.work/api/user/checkin"
	cookie := os.Getenv("GLADOS_COOKIE")
	body := sendRequest("POST", baseUrl, cookie)
	reg := regexp.MustCompile(`.*"message":"([^"]*)".*`)
	results := reg.FindStringSubmatch(body)
	if len(results) > 0 {
		message := results[1]
		fmt.Println(message)
		return true
	}
	return false
}

func sign() {
	return
}

func main() {
}
