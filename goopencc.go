package goopencc

import (
	"bytes"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	ZH2TW = "zhs2zhtw_p.ini"
	TW2ZH = "zhtw2zhcn_s.ini"
)

func Zh2Tw(v string) (string, int) {
	v = strings.Trim(v, " ")
	if v == "" {
		return v, 200
	}
	return translate(v, ZH2TW)
}

func Tw2Zh(v string) (string, int) {
	v = strings.Trim(v, " ")
	if v == "" {
		return v, 200
	}
	return translate(v, TW2ZH)
}

func translate(v string, m string) (string, int) {
	apiUrl := "http://opencc.byvoid.com/convert"
	data := url.Values{}
	data.Set("text", v)
	data.Add("config", m)
	data.Add("precise", "0")

	client := &http.Client{}
	r, _ := http.NewRequest("POST", apiUrl, bytes.NewBuffer([]byte(data.Encode())))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, _ := client.Do(r)
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return buf.String(), resp.StatusCode
}
