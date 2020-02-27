package main

import (
	"context"
	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/workflow"
	"time"
)

func init() {
	workflow.Register(SampleWorkflow)
	workflow.Register(SampleLocalWorkflow)

	activity.Register(sampleActivity)
}

// Sample structure for activity result
type result = struct{}

// Sample Workflow without local activity
func SampleWorkflow(ctx workflow.Context, nilResult bool) (r *result, err error) {
	var (
		ao = workflow.ActivityOptions{
			ScheduleToStartTimeout: 5 * time.Second,
			StartToCloseTimeout:    time.Minute,
		}
		ctx1 = workflow.WithActivityOptions(ctx, ao)
	)

	err = workflow.ExecuteActivity(ctx1, sampleActivity, nilResult).Get(ctx1, &r)

	return
}

// Sample Workflow with local activity
func SampleLocalWorkflow(ctx workflow.Context, nilResult bool) (r *result, err error) {
	var (
		ao = workflow.LocalActivityOptions{
			ScheduleToCloseTimeout: 5 * time.Second,
		}
		ctx1 = workflow.WithLocalActivityOptions(ctx, ao)
	)

	err = workflow.ExecuteLocalActivity(ctx1, sampleActivity, nilResult).Get(ctx1, &r)

	return
}

func sampleActivity(ctx context.Context, nilResult bool) (*result, error) {
	if nilResult {
		return nil, nil
	}
	return &result{}, nil
}
