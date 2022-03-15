package cmd

import (
	"fmt"
	"strconv"
	"time"

	aw "github.com/deanishe/awgo"
	"github.com/spf13/cobra"
)

var (
	iconClock = &aw.Icon{
		Value: aw.IconClock.Value,
		Type:  aw.IconClock.Type,
	}

	layouts = []string{
		"2006-01-02 15:04:05",
		"2006-01-02T15:04:05Z",
	}

	moreLayouts = []string{
		"2006-01-02",
		"2006-01-02 15:04",
		"2006-01-02 15:04:05",
	}
)

var dateCmd = &cobra.Command{
	Use: "date",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			return
		}

		defer func() {
			wf.SendFeedback()
			return
		}()

		getNowTime(args)   // 获取当前时间
		getByTimeStr(args) // 根据时间字符串获取时间
	},
}

// 获取当前时间-三种格式
func getNowTime(args []string) {
	if args[0] != "now" {
		return
	}

	now := time.Now()
	secs := fmt.Sprintf("%d", now.Unix())
	wf.NewItem(secs).
		Subtitle("unix timestamp").
		Icon(iconClock).
		Arg(secs).
		Valid(true)

	// process all time layouts
	for _, layout := range layouts {
		v := now.Format(layout)
		wf.NewItem(v).
			Subtitle(layout).
			Icon(iconClock).
			Arg(v).
			Valid(true)
	}
}

type layoutStruct struct {
	layout string
	Time   time.Time
}

// 根据时间字符串获取时间
func getByTimeStr(args []string) {
	if args[0] == "now" {
		return
	}

	timeStr := args[0]
	matchLayout := func(layouts []string, timeStr string) *layoutStruct {
		for _, layout := range layouts {
			t, err := time.ParseInLocation(layout, timeStr, time.Local)
			if err == nil {
				return &layoutStruct{
					layout: layout,
					Time:   t,
				}
			}
		}
		return nil
	}

	l := matchLayout(layouts, timeStr)
	if l == nil { // layouts样式不匹配
		l = matchLayout(moreLayouts, timeStr)
		if l == nil { // moreLayouts样式不匹配
			parseInt, err := strconv.ParseInt(timeStr, 10, 64)
			if err != nil {
				return
			}
			for _, layout := range layouts {
				v := time.Unix(parseInt, 0).Format(layout)
				wf.NewItem(v).
					Subtitle(layout).
					Icon(iconClock).
					Arg(v).
					Valid(true)
			}
			return
		}
	}

	// prepend unix timeObj
	secs := fmt.Sprintf("%d", l.Time.Unix())
	wf.NewItem(secs).
		Subtitle("unix timestamp").
		Icon(iconClock).
		Arg(secs).
		Valid(true)

	// other time layouts
	for _, layout := range layouts {
		if layout == l.layout {
			continue
		}
		v := l.Time.Format(layout)
		wf.NewItem(v).
			Subtitle(layout).
			Icon(iconClock).
			Arg(v).
			Valid(true)
	}

	return
}
