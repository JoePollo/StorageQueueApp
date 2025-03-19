package main

import (
	"context"
	"fmt"
	"log"
	// "os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azqueue"
)

func ErrorHandler(object string, err error) {
	log.Fatal(fmt.Errorf("Failed to build %s due to error:\n%w", object, err))
}

func main() {

	// Env := os.Getenv("ENV")

	creds, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		ErrorHandler("NewDefaultAzureCredential", err)
	}

	client, err := azqueue.NewQueueClient("https://styymdevuscecafaf3f9.queue.core.windows.net/stq-yym-dev-usce",
		creds,
		nil)
	if err != nil {
		ErrorHandler("NewClient", err)
	}

	response, err := client.EnqueueMessage(context.TODO(), "super cool message", nil)

	if err != nil {
		ErrorHandler("EnqueueMessage", err)
	}

	fmt.Printf("Message ID:\n%v\n", *response.Messages[0].MessageID)

	message, err := client.DequeueMessage(context.TODO(), nil)
	if err != nil {
		ErrorHandler("DequeueMessage", err)
	}

	fmt.Printf("Dequeued message:\n%v\n", *message.Messages[0].MessageText)

}
