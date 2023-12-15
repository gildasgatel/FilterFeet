package main

import (
	"filterfleet/pkg/filter"
	"filterfleet/pkg/task"
	"filterfleet/pkg/utils"
	"fmt"
	"runtime"
	"time"

	"flag"
	"log"
	"os"
)

func main() {
	t := time.Now()
	// Setup CLI commands
	srcInput := flag.String("src", "imgs", "Source directory containing the images")
	dstInput := flag.String("dst", "output", "Destination directory to store the output images")
	filterInput := flag.String("filter", "grayscale", "Choose the filter to apply (grayscale or blur)")
	taskInput := flag.String("task", "waitgroup", "Specify the task type (channel or waygroup).")
	poolsizeInput := flag.Int("poolsize", 2, "Set the CPU pool size ")
	flag.Parse()

	err := utils.CreateDirIfNotExist(*dstInput)
	if err != nil {
		log.Fatal(err)
	}
	runtime.GOMAXPROCS(*poolsizeInput)

	filesEntry, err := os.ReadDir(*srcInput)
	if err != nil {
		log.Fatal(err)
	}
	// init filter
	fil, err := filter.New(*filterInput)
	if err != nil {
		log.Fatal(err)
	}
	// set blur sigma
	fil.SetSigma(20)

	// Init tasker
	app := task.New(fil, filesEntry)
	app.SetDst(*dstInput)
	app.SetSrc(*srcInput)

	switch *taskInput {
	case "waitgroup":
		app.ProcessWg()
	case "channel":
		app.ProcessChan()
	}
	duration := time.Since(t)
	fmt.Println("Process time: " + duration.String())

}
