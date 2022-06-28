package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"google.golang.org/api/cloudbilling/v1"
	"google.golang.org/api/googleapi"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	project_b_acc := os.Getenv("GCP_PROJECT_B_ACC")
	project := os.Getenv("GCP_PROJECT_ID")
	if project == "" {
		log.Panic("PROJECT_ID empty")
		return
	}

	p := cloudbilling.ProjectBillingInfo{
		BillingAccountName: project_b_acc,
		BillingEnabled:     false,
		Name:               project,
		ProjectId:          project,
		ServerResponse:     googleapi.ServerResponse{},
		ForceSendFields:    nil,
		NullFields:         nil,
	}

	cloudbillingService, err := cloudbilling.NewService(ctx)
	if err != nil {
		log.Panic("problem with billingService")
	}

	b := cloudbillingService.Projects.UpdateBillingInfo(project, &p)
	fmt.Printf("%s\n", b)

	log.Println(cloudbillingService.BasePath)

	client, err := pubsub.NewClient(ctx, project)
	if err != nil {
		// TODO: Handle error.
	}

	topicId := os.Getenv("GCP_PUBSUB_TOPIC")
	topic := client.Topic(topicId)
	defer topic.Stop()
	var results []*pubsub.PublishResult
	r := topic.Publish(ctx, &pubsub.Message{
		Data: []byte("hello world"),
	})
	results = append(results, r)
	// Do other work ...
	for _, r := range results {
		id, err := r.Get(ctx)
		if err != nil {
			// TODO: Handle error.
		}
		fmt.Printf("Published a message with a message ID: %s\n", id)
	}
}
