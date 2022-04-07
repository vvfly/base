package utils

import (
	"fmt"
	"time"
)

// WeekByDate 判断时间是当年的第几周
func WeekByDate(t time.Time) string {

	yearDay := t.YearDay()
	yearFirstDay := t.AddDate(0, 0, -yearDay+1)

	// 1月1日是周几
	firstDayInWeek := int(yearFirstDay.Weekday())
	// 第一周有几天
	firstWeekDays := 1
	if firstDayInWeek != 0 {
		firstWeekDays = 7 - firstDayInWeek + 1
	}

	var week int
	if yearDay <= firstWeekDays {
		week = 1
	} else {
		week = (yearDay-firstWeekDays)/7 + 2
	}

	return fmt.Sprintf("%d第%d周", t.Year(), week)
}
