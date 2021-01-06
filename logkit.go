package kits2

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

const (
	sIZE      = 500  //  log存储空间总长度
	fetchSIZE = 40   //  一次拿多少
	strSIZE   = 3000 //  单条日志的最大限制
)

type LogKit struct {
	logs   *CircleQueue
	readme string //  名称的注释
}

func NewLogKit(readme string) *LogKit {
	lk := &LogKit{}
	lk.readme = readme
	lk.logs = NewCircleQueue(sIZE)
	return lk
}

//  展示数据
func (lk *LogKit) Show() string {
	str := lk.titleStr() + formatFetchedLog(lk.logs.Get(fetchSIZE))
	return str
}

func (lk *LogKit) titleStr() string {
	str := ""
	str += "\n----------------------\n " + "   \n" + "日志说明:" + lk.readme + "\n"
	return str
}

//  把日志信息放入
func (lk *LogKit) Put(a ...string) {
	lk.logs.Put(formatLogForPut(a...))
}

// 输入的日志内容，增加额外的信息以及处理，如时间信息，或者截断等
func formatLogForPut(strs ...string) string {
	buffer := bytes.Buffer{}
	now := time.Now()
	buffer.WriteString(now.Format("2006-01-02 15:04:05") + " 0." + fmt.Sprint(now.Nanosecond()))
	buffer.WriteString("\n")
	for _, v := range strs {
		str := fmt.Sprint(v)
		if len(str) > strSIZE {
			str = str[0:strSIZE] + "......后面的内容过长截断......"
		}
		buffer.WriteString(str)
		//  每一个参数后面都加一个换行
		buffer.WriteString("\n")
	}
	buffer.WriteString("\n\n")
	return buffer.String()
}

//  美化输出
func formatFetchedLog(values []string, id uint64) string {
	buffer := bytes.Buffer{}
	buffer.WriteString("序号: " + strconv.FormatUint(id, 10) + "\n")
	for _, v := range values {
		buffer.WriteString(v)
	}
	return buffer.String()
}
