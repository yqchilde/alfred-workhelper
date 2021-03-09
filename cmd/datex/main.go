package main

import (
	"github.com/deanishe/awgo"

	"alfred/internal/service"
)

// Workflow is the main API
var (
	wf *aw.Workflow
)

func main() {
	wf = aw.New()
	wf.Run(service.RunDateX)
}
