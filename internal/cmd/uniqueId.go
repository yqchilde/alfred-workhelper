package cmd

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var uniqueIdCmd = &cobra.Command{
	Use: "id",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			return
		}

		defer func() {
			wf.SendFeedback()
			return
		}()

		generateUUID(args)
		snowflakeID(args)
	},
}

// generateUUID generates a UUID
func generateUUID(args []string) {
	if args[0] != "get" {
		return
	}

	newUUID, err := uuid.NewUUID()
	if err != nil {
		wf.FatalError(err)
	}

	uuid36 := newUUID.String()
	wf.NewItem(uuid36).
		Subtitle("36位UUID").
		Arg(uuid36).
		Valid(true)

	uuid32 := strings.Replace(uuid36, "-", "", -1)
	wf.NewItem(uuid32).
		Subtitle("32位UUID").
		Arg(uuid32).
		Valid(true)
}

// snowflakeID generates a snowflake ID
func snowflakeID(args []string) {
	if args[0] != "get" {
		return
	}

	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Generate a snowflake ID.
	id := node.Generate()
	wf.NewItem(id.String()).
		Subtitle("雪花算法").
		Arg(id.String()).
		Valid(true)
}
