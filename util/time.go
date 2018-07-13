package util

import (
	"time"
)

// StandardTimeFormat 标准格式
var StandardTimeFormat = "2006-01-02 15:04:05"

// TimestampToTime 时间戳到时间
// timestamp int64:转换的时间戳
// layout string 需要转换的时间格式
func TimestampToTime(timestamp int64, layout string) string {
	return time.Unix(timestamp, 0).Format(layout)
}

// TimeStrToTime 时间字符串到时间，你需要知道你的时间字符串格式
// layout 解析的时间格式，
// timestr 转换的时间
func TimeStrToTime(layout, timestr string) (ttime time.Time, err error) {
	var (
		loc *time.Location
	)
	loc, err = time.LoadLocation("Local")
	if err != nil {
		return
	}
	ttime, err = time.ParseInLocation(layout, timestr, loc)
	if err != nil {
		return
	}
	return
}
