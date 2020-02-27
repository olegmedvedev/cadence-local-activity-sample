package main

import (
	"flag"
	"github.com/pborman/uuid"
	"github.com/uber-common/cadence-samples/cmd/samples/common"
	"go.uber.org/cadence/client"
	"go.uber.org/cadence/worker"
	"time"
)

const taskList = "MainTaskList"

func startWorkers(h *common.SampleHelper) {
	workerOptions := worker.Options{
		MetricsScope: h.Scope,
		Logger:       h.Logger,
	}
	h.StartWorkers(h.Config.DomainName, taskList, workerOptions)
}

func startWorkflow(h *common.SampleHelper, local, nilRes bool) {
	var workflowOptions = client.StartWorkflowOptions{
		TaskList:                        taskList,
		ExecutionStartToCloseTimeout:    time.Hour,
		DecisionTaskStartToCloseTimeout: time.Hour,
		ID:                              uuid.New(),
	}

	if local {
		h.StartWorkflow(workflowOptions, SampleLocalWorkflow, nilRes)
	} else {
		h.StartWorkflow(workflowOptions, SampleWorkflow, nilRes)
	}
}

func main() {
	var trigger, nilRes, local bool
	flag.BoolVar(&trigger, "t", false, "Trigger starting workflow")
	flag.BoolVar(&nilRes, "n", false, "Return nil in activity")
	flag.BoolVar(&local, "l", false, "Execute with local activity")

	flag.Parse()

	var h common.SampleHelper
	h.SetupServiceConfig()

	if trigger {
		startWorkflow(&h, local, nilRes)
	} else {
		startWorkers(&h)
		select {}
	}
}
