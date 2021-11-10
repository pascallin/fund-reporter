package datasource

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
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
	GDP = "gdp"
	CPI = "cpi"
	PPI = "ppi"
)

func GetEconomicStat(statType string, startTime string, endTime string) (*Res, error) {
	url, err := getUrl(statType, startTime, endTime)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	return getStatAPIData(client, url)
}

func getUrl(statType string, startTime string, endTime string) (string, error) {
	const (
		gdpURL = "https://bmfw.www.gov.cn/bjww/StatisSelectRedis/GDPlist.do"
		cpiURL = "https://bmfw.www.gov.cn/bjww/StatisSelectRedis/CPIlist.do"
		ppiURL = "https://bmfw.www.gov.cn/bjww/StatisSelectRedis/PPIlist.do"
	)
	params := url.Values{}
	params.Add("qinStart", startTime)
	params.Add("qinEnd", endTime)
	var baseUrl string
	var qinName string
	switch statType {
	case GDP:
		baseUrl = gdpURL
		qinName = "国内生产总值（亿元）"
	case CPI:
		baseUrl = cpiURL
		qinName = "居民消费价格月度同比涨跌（%）"
	case PPI:
		baseUrl = ppiURL
		qinName = "工业生产者出厂价格月度同比涨跌（%）"
	default:
		return "", errors.New("wrong stat type")
	}
	params.Add("qinName", qinName)
	requestUrl, err := url.Parse(baseUrl)
	if err != nil {
		log.Error(err)
		return "", nil
	}
	requestUrl.RawQuery = params.Encode() // Escape Query Parameters
	log.WithField("Encoded URL", requestUrl.String()).Info()
	return requestUrl.String(), nil
}

func getSignature(time int64) string {
	h := sha256.New()
	secret := fmt.Sprintf("%dfTN2pfuisxTavbTuYVSsNJHetwq5bJvCQkjjtiLM2dCratiA%d", int(time), int(time))
	h.Write([]byte(secret))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

func getStatAPIData(client *http.Client, fullUrl string) (*Res, error) {
	now := time.Now()
	sec := now.Unix()
	signature := getSignature(sec)

	log.WithFields(log.Fields{
		"signature": signature,
		"timestamp": sec,
	}).Info("getSignature")

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("x-wif-nonce", "QkjjtiLM2dCratiA")
	req.Header.Add("x-wif-paasid", "smt-application")
	req.Header.Add("x-wif-signature", signature)
	req.Header.Add("x-wif-timestamp", fmt.Sprintf("%d", sec))
	req.Header.Add("sec-ch-ua", `Google Chrome";v="95", "Chromium";v="95", ";Not A Brand";v="99"`)
	req.Header.Add("sec-ch-ua-platform", "Windows")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("Host", "bmfw.www.gov.cn")
	req.Header.Add("Origin", "http://www.gov.cn/")
	req.Header.Add("Referer", "http://www.gov.cn/")
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	log.WithField("StatusCode", resp.StatusCode).Info("HTTP Response")

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
		log.Error("Can not decode JSON")
		return nil, err
	}
	return &decodeResult, nil
}
