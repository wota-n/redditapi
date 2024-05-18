package main

import (
    "github.com/vartanbeno/go-reddit/v2/reddit"
    "log"
	"context"
	"fmt"
)

func main() {
	credentials := reddit.Credentials{
		ID:       "",
		Secret:   "",
		Username: "",
		Password: "",
	}

    client, err := reddit.NewClient(credentials)
    if err != nil {
        log.Fatalf("failed to create client: %v", err)
    }

	// Assuming 'client' is already authenticated
	comments, _, err := client.User.Comments(context.Background(), &reddit.ListUserOverviewOptions{
		ListOptions: reddit.ListOptions{
			Limit: 1,
		},
	})
	if err != nil {
		log.Fatalf("failed to fetch comments: %v", err)
	}

	for _, comment := range comments {
		fmt.Println(comment.Body)
		thingID := "t1_" + comment.ID
		_, _, err := client.Comment.Edit(context.Background(), thingID, test)
		
		if err != nil {
			log.Fatalf("failed to edit comment: %v", err)
		}
	}



}
