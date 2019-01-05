package utils

import (
	"strconv"
	"time"
)

const baseFormat = "2006-01-02 15:04:05"

type TimeMove struct {
	Year    int
	Month   int
	Day     int
	Hour    int
	Minute  int
	Seconds int
}

// 获取当前时间时间戳
func GetCurTimeStamp() int64 {
	curTime := time.Now().Unix()
	return curTime
}

// 获取当前时间(str类型)
func GetCurTimeStr() string {
	curTime := time.Now().Format(baseFormat)
	return curTime
}

// 获取当前时间(time.Time类型)
func GetCurTimeUtc() time.Time {
	curTime := time.Now().UTC()
	return curTime
}

// 获取当前日期(str类型)
func GetCurDate() string {
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()
	dayStr := strconv.Itoa(day)
	curDate := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-"
	if len(dayStr) < 2 {
		dayStr = "0" + dayStr
	}
	curDate = curDate + dayStr
	return curDate
}

// str时间转time.Time时间
func StrToDateTime(strTime string) (time.Time, error) {
	parseStrTime, err := time.Parse(baseFormat, strTime)
	if err != nil {
		return time.Time{}, err
	}
	return parseStrTime, nil
}

// 获取输入时间相隔一定时间的时间
func GetAnotherTime(beforeTime time.Time, timeMove TimeMove) time.Time {
	afterTime := beforeTime.AddDate(timeMove.Year, timeMove.Month, timeMove.Day)
	hour := time.Duration(timeMove.Hour)
	minute := time.Duration(timeMove.Minute)
	seconds := time.Duration(timeMove.Seconds)
	date := afterTime.Add(hour*time.Hour + minute*time.Minute + seconds*time.Second)
	return date
}
