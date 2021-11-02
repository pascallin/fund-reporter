package datasource

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ResData struct {
	Datetime string   `json:"datetime"`
	Values   []string `json:"values"`
}
type Res struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data []ResData `json:"data"`
}

const (
	GDP = iota
	CPI
	PPI
)
const (
	gdpURL = "https://bmfw.www.gov.cn/bjww/StatisSelectRedis/GDPlist.do?qinStart=2012-A&qinEnd=2021-C&qinName=%E5%9B%BD%E5%86%85%E7%94%9F%E4%BA%A7%E6%80%BB%E5%80%BC%EF%BC%88%E4%BA%BF%E5%85%83%EF%BC%89"
	cpiURL = "https://bmfw.www.gov.cn/bjww/StatisSelectRedis/CPIlist.do?qinStart=2016-01&qinEnd=2021-09&qinName=%E5%B1%85%E6%B0%91%E6%B6%88%E8%B4%B9%E4%BB%B7%E6%A0%BC%E6%9C%88%E5%BA%A6%E5%90%8C%E6%AF%94%E6%B6%A8%E8%B7%8C%EF%BC%88%25%EF%BC%89"
	ppiURL = "https://bmfw.www.gov.cn/bjww/StatisSelectRedis/PPIlist.do?qinStart=2010-01&qinEnd=2021-09&qinName=%E5%B7%A5%E4%B8%9A%E7%94%9F%E4%BA%A7%E8%80%85%E5%87%BA%E5%8E%82%E4%BB%B7%E6%A0%BC%E6%9C%88%E5%BA%A6%E5%90%8C%E6%AF%94%E6%B6%A8%E8%B7%8C%EF%BC%88%25%EF%BC%89"
)

func GetEconomicStat(statType int) (*Res, error) {
	var url string
	switch statType {
	case GDP:
		url = gdpURL
	case CPI:
		url = cpiURL
	case PPI:
		url = ppiURL
	default:
		return nil, errors.New("wrong stat type")
	}
	client := &http.Client{}
	return GetStatAPIData(client, url)
}

func GetStatAPIData(client *http.Client, fullUrl string) (*Res, error) {
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("x-wif-nonce", "QkjjtiLM2dCratiA")
	req.Header.Add("x-wif-paasid", "smt-application")
	req.Header.Add("x-wif-signature", "697A55043A808898AC95260A40945999DE8C7D95ADE66F74A654C5895C928EDF")
	req.Header.Add("x-wif-timestamp", "1635835843")
	req.Header.Add("sec-ch-ua", `Google Chrome";v="95", "Chromium";v="95", ";Not A Brand";v="99"`)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	fmt.Printf("StatusCode: %v\n", resp.StatusCode)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	// using Unmarshal
	// body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(body)) // convert to string before print
	// var unmarshalResult GDPRes
	// if err := json.Unmarshal(body, &unmarshalResult); err != nil { // Parse []byte to the go struct pointer
	// 	fmt.Println("Can not unmarshal JSON")
	// }
	// fmt.Printf("Results: %v\n", unmarshalResult)

	// using decoder
	var decodeResult Res
	decoder := json.NewDecoder(resp.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&decodeResult)
	if err != nil {
		fmt.Println("Can not decode JSON")
		return nil, err
	}
	return &decodeResult, nil
}
