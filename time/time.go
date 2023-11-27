package main

import (
	"fmt"
	"time"
)

func main() {
	timeobj := time.Now()

	// fmt.Println(timeobj)

	// year := timeobj.Year()
	// month := timeobj.Month()
	// day := timeobj.Day()
	// hour := timeobj.Hour()
	// minute := timeobj.Minute()
	// second := timeobj.Second()
	// fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minute, second)

	//使用format来格式化,按照以下格式来
	//2006  年
	// 01  月
	// 02  日
	// 03  时  03是12小时制，15是24小时制
	// 04  分
	// 05  秒
	// var str = timeobj.Format("2006/01/02 15:04:05")
	// var str = timeobj.Format("2006/01/02 03:04:05")
	// fmt.Println(str)

	unixtime := timeobj.Unix()
	fmt.Println(unixtime) //毫秒
	unixNAtime := timeobj.UnixNano()
	fmt.Println(unixNAtime) //纳秒

}
