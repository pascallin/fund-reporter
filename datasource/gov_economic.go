package datasource

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var gdpURL = "https://bmfw.www.gov.cn/bjww/StatisSelectRedis/GDPlist.do?qinStart=2012-A&qinEnd=2021-C&qinName=%E5%9B%BD%E5%86%85%E7%94%9F%E4%BA%A7%E6%80%BB%E5%80%BC%EF%BC%88%E4%BA%BF%E5%85%83%EF%BC%89"
var cpiURL = "https://bmfw.www.gov.cn/bjww/StatisSelectRedis/CPIlist.do?qinStart=2016-01&qinEnd=2021-09&qinName=%E5%B1%85%E6%B0%91%E6%B6%88%E8%B4%B9%E4%BB%B7%E6%A0%BC%E6%9C%88%E5%BA%A6%E5%90%8C%E6%AF%94%E6%B6%A8%E8%B7%8C%EF%BC%88%25%EF%BC%89"
var ppiURL = "https://bmfw.www.gov.cn/bjww/StatisSelectRedis/PPIlist.do?qinStart=2010-01&qinEnd=2021-09&qinName=%E5%B7%A5%E4%B8%9A%E7%94%9F%E4%BA%A7%E8%80%85%E5%87%BA%E5%8E%82%E4%BB%B7%E6%A0%BC%E6%9C%88%E5%BA%A6%E5%90%8C%E6%AF%94%E6%B6%A8%E8%B7%8C%EF%BC%88%25%EF%BC%89"

type GDPResData struct {
	Datetime string   `json:"datetime"`
	Values   []string `json:"values"`
}
type GDPRes struct {
	Code int64        `json:"code"`
	Msg  string       `json:"msg"`
	Data []GDPResData `json:"data"`
}

const (
	GDP = iota
	CPI
	PPI
)

func GetEconomicStat(statType int) {
	/**
	GET /bjww/StatisSelectRedis/GDPlist.do?qinStart=2012-A&qinEnd=2021-C&qinName=%E5%9B%BD%E5%86%85%E7%94%9F%E4%BA%A7%E6%80%BB%E5%80%BC%EF%BC%88%E4%BA%BF%E5%85%83%EF%BC%89 HTTP/1.1
	Host: bmfw.www.gov.cn
	Connection: keep-alive
	sec-ch-ua: "Google Chrome";v="95", "Chromium";v="95", ";Not A Brand";v="99"
	x-wif-nonce: QkjjtiLM2dCratiA
	sec-ch-ua-mobile: ?0
	User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.54 Safari/537.36
	x-wif-paasid: smt-application
	Content-Type: application/json; charset=utf-8
	Access-Control-Allow-Origin: *
	x-wif-signature: 697A55043A808898AC95260A40945999DE8C7D95ADE66F74A654C5895C928EDF
	x-wif-timestamp: 1635835843
	sec-ch-ua-platform: "Windows"
	Origin: http://www.gov.cn
	Sec-Fetch-Site: cross-site
	Sec-Fetch-Mode: cors
	Sec-Fetch-Dest: empty
	Referer: http://www.gov.cn/
	Accept-Encoding: gzip, deflate, br
	Accept-Language: zh-CN,zh;q=0.9
	*/
	var url string
	switch statType {
	case GDP:
		url = gdpURL
	case CPI:
		url = cpiURL
	case PPI:
		url = ppiURL
	default:
		panic("wrong stat type")
	}
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("x-wif-nonce", "QkjjtiLM2dCratiA")
	request.Header.Add("x-wif-paasid", "smt-application")
	request.Header.Add("x-wif-signature", "697A55043A808898AC95260A40945999DE8C7D95ADE66F74A654C5895C928EDF")
	request.Header.Add("x-wif-timestamp", "1635835843")
	request.Header.Add("sec-ch-ua", `Google Chrome";v="95", "Chromium";v="95", ";Not A Brand";v="99"`)
	request.Header.Add("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("StatusCode: %v\n", resp.StatusCode)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	// using decoder
	var decodeResult GDPRes
	decoder := json.NewDecoder(resp.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&decodeResult)
	if err != nil {
		fmt.Println("Can not decode JSON")
	}
	fmt.Printf("Results: %v\n", decodeResult)

	// using Unmarshal
	// body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	// if err != nil {
	// 	panic(err)
	// }
	// // fmt.Println(string(body)) // convert to string before print
	// var unmarshalResult GDPRes
	// if err := json.Unmarshal(body, &unmarshalResult); err != nil { // Parse []byte to the go struct pointer
	// 	fmt.Println("Can not unmarshal JSON")
	// }

	// fmt.Printf("Results: %v\n", unmarshalResult)
}
