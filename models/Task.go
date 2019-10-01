package models

import (
	"fmt"
	"fortuneway/common"
	"github.com/Tang-RoseChild/mahonia"
	"strconv"
	"github.com/astaxie/beego/toolbox"
	"io/ioutil"
	"net/http"
	"strings"
)

var ticketCode = "sz000815"

var initPrice = getPrice()
var initVolume = getVolume()

func Task() {
	lastPrice := initPrice

	newPriceMonitorTask := toolbox.NewTask("newPriceMonitor", "0/1 * 9-22 * * 1-5", func() error {
		newPrice := getPrice()

		rate := calculateRate(lastPrice, newPrice)

		if rate > 3 {
			// 发送短信提示
		}

		fmt.Println("rate: ", rate)

		return nil
	})

	min5VolumeMonitorTask := toolbox.NewTask("volumeMonitor", "0/1 * 9-22 * * 1-5", func() error {

		return nil
	})

	err := newPriceMonitorTask.Run()
	if err != nil {
		fmt.Println(err)
	}

	toolbox.AddTask(newPriceMonitorTask.Taskname, newPriceMonitorTask)
	toolbox.AddTask(min5VolumeMonitorTask.Taskname, min5VolumeMonitorTask)
	toolbox.StartTask()
	defer toolbox.StopTask()

	//beego.Run()
}

func getPrice() int {
	resp, err := http.Get(common.Ontime_Price_Address + ticketCode)
	if err != nil {
		println("query failed")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println("transfer body failed")
	}

	originalStr := mahonia.NewDecoder("gbk").ConvertString(string(body))
	dataStr := originalStr[strings.Index(originalStr, "\"") + 1 : strings.LastIndex(originalStr, "\"")]
	dataArr := strings.Split(dataStr, ",")

	println("campany_name: " + dataArr[0], "price: " + dataArr[3])
	price, err := strconv.Atoi(dataArr[3])

	return price
}

func getVolume() int {
	//resp, err := http.Get(common.Min_Volume_Address + ticketCode)
	//if err != nil {
	//	println("query failed")
	//}
	//defer resp.Body.Close()
	//
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	println("transfer body failed")
	//}
	//
	//originalStr := mahonia.NewDecoder("gbk").ConvertString(string(body))
	//
	//
	//println("campany_name: " + dataArr[0], "price: " + dataArr[3])
	//price, err := strconv.Atoi(dataArr[3])

	return 0
}

// 计算涨跌幅
func calculateRate(lastData int, newData int) int64 {
	return 0.0
}
