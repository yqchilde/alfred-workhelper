package cmd

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
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

	regexpTimestamp = regexp.MustCompile(`^[1-9]\d+$`)
)

var dateCmd = &cobra.Command{
	Use:   "date",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if len(args) == 0 {
			return
		}

		defer func() {
			if err == nil {
				wf.SendFeedback()
				return
			}
		}()

		// now
		input := strings.Join(args, " ")
		if input == "now" {
			processNow()
			return
		}

		// timestamp
		if regexpTimestamp.MatchString(input) {
			v, e := strconv.ParseInt(args[0], 10, 32)
			if e == nil {
				processTimestamp(time.Unix(v, 0))
				return
			}
			err = e
			return
		}

		// time string
		err = processTimeStr(input)
	},
}

// process the current time
func processNow() {
	now := time.Now()

	secs := fmt.Sprintf("%d", now.Unix())
	wf.NewItem(secs).
		Subtitle("unix timestamp").
		Icon(iconClock).
		Arg(secs).
		Valid(true)

	// process all time layouts
	processTimestamp(now)
}

// process all time layouts
func processTimestamp(timestamp time.Time) {
	for _, layout := range layouts {
		v := timestamp.Format(layout)
		wf.NewItem(v).
			Subtitle(layout).
			Icon(iconClock).
			Arg(v).
			Valid(true)
	}
}

func processTimeStr(timeStr string) error {

	timeObj := time.Time{}
	layoutMatch := ""

	layoutMatch, timeObj, ok := matchedLayout(layouts, timeStr)
	if !ok {
		layoutMatch, timeObj, ok = matchedLayout(moreLayouts, timeStr)
		if !ok {
			return errors.New("no matched time layout found")
		}
	}

	// prepend unix timeObj
	secs := fmt.Sprintf("%d", timeObj.Unix())
	wf.NewItem(secs).
		Subtitle("unix timestamp").
		Icon(iconClock).
		Arg(secs).
		Valid(true)

	// other time layouts
	for _, layout := range layouts {
		if layout == layoutMatch {
			continue
		}
		v := timeObj.Format(layout)
		wf.NewItem(v).
			Subtitle(layout).
			Icon(iconClock).
			Arg(v).
			Valid(true)
	}

	return nil
}

func matchedLayout(layouts []string, timeStr string) (string, time.Time, bool) {
	loc, _ := time.LoadLocation("Local")
	for _, layout := range layouts {
		time, err := time.ParseInLocation(layout, timeStr, loc)
		if err == nil {
			return layout, time, true
		}
	}

	return "", time.Time{}, false
}
