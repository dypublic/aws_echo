package main

import (
	//"encoding/json"
	"io/ioutil"
	"net/http"
	//"os"
	"time"
    "fmt"
	"encoding/json"
)

/*
http://169.254.169.254/latest/meta-data/
hostname
instance-id
local-hostname
local-ipv4
placement/
profile
public-hostname
public-ipv4
*/
var (
	awsMetaDataUrl = "http://169.254.169.254/latest/meta-data/"
	awsZone        = "placement/availability-zone"
	awsIpv4        = "local-ipv4"
	awsId          = "instance-id"
)

type AWSMetaData struct {
	Zone string `json:"availability-zone"`
	Ip   string `json:"local-ipv4"`
	Id   string `json:"instance-id"`
}

func httpGet(path string) (string, error) {
	client := http.Client{
		Timeout: time.Second,
	}
	resp, err := client.Get(path)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func GetMetaData() AWSMetaData {
	data := map[string]string{}
	item := []string{awsZone, awsIpv4, awsId}
	for _, path := range item {
		value, err := httpGet(awsMetaDataUrl + path)
		if err != nil {
            fmt.Println("Get ", awsMetaDataUrl+path, " error:", err)
			continue
		}
		data[path] = value
	}
	result := AWSMetaData{Zone: data[awsZone], Id: data[awsId], Ip: data[awsIpv4]}
	fmt.Println("aws meta data:", result)
	return result
}

func handler(w http.ResponseWriter, r *http.Request) {
    data := GetMetaData()
    res, _ := json.Marshal(data)
    w.Write(res)
    //fmt.Fprintf(w, "%v\n", result)
}

var StaticRes []byte

func init() {
	data := GetMetaData()
	StaticRes, _ = json.Marshal(data)
}

func handlerStatic(w http.ResponseWriter, r *http.Request) {
	w.Write(StaticRes)
}

func main() {
    http.HandleFunc("/", handler)
	http.HandleFunc("/static", handlerStatic)
    http.ListenAndServe(":8080", nil)
}