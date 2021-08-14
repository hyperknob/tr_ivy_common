package tr_ivy_common

import (
	"strings"
	"time"
)

type JsonTime time.Time

const timeFormat = "2006-01-02 15:04:05"

func (t *JsonTime) UnmarshalJSON(data []byte) (err error) {
	dataStr := string(data)
	var now time.Time
	// 以下语句为了兼容RFC3339格式
	dataStr = strings.ReplaceAll(dataStr, "T", " ")
	dataStr = strings.ReplaceAll(dataStr, "+08:00", "")
	now, err = time.ParseInLocation(`"`+timeFormat+`"`, dataStr, time.Local)
	*t = JsonTime(now)
	return
}

func (t JsonTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}

func (t JsonTime) String() string {
	return time.Time(t).Format(timeFormat)
}