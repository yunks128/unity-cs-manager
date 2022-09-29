package actions

import (
	"context"
	"fmt"
	"github.com/google/go-github/v47/github"
	"log"
)

func Execute(meta string) {
	client := github.NewClient(nil)
	req := github.CreateWorkflowDispatchEventRequest{
		Ref:    "",
		Inputs: nil,
	}
	wf, err := client.Actions.CreateWorkflowDispatchEventByID(context.Background(), "unity-sds", "unity-cs-infra", int64(2979181266), req)

	if err != nil {
		log.Fatalf("Error triggering action: %v", err)
	}
	fmt.Println(wf)
}

func TearDown(projname string) {

}

func List() {

}
