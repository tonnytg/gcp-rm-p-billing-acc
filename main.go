package main

import (
	"context"
	"google.golang.org/api/cloudbilling/v1"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	p := cloudbilling.ProjectBillingInfo{
		BillingAccountName: "",
		BillingEnabled:     false,
		Name:               "",
		ProjectId:          "",
		ServerResponse:     googleapi.ServerResponse{},
		ForceSendFields:    nil,
		NullFields:         nil,
	}

	cloudbillingService, err := cloudbilling.NewService(ctx, option.WithAPIKey(os.Getenv("GCP_API_TOKEN")))
	if err != nil {
		log.Fatal(err)
	}

	b, err := cloudbillingService.Projects.UpdateBillingInfo("projects/remove-billing-acc", &p).Do()
	if err != nil {
		log.Panic(err)
	}
	log.Println(b)
}
