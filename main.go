package main

import (
    "github.com/vartanbeno/go-reddit/v2/reddit"
    "log"
	"context"
	"fmt"
	"github.com/jaswdr/faker"
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

	randomParagraph := fmt.Sprintf(generateRandomParagraph())

	for _, comment := range comments {
		fmt.Println(comment.Body)
		thingID := "t1_" + comment.ID
		_, _, err := client.Comment.Edit(context.Background(), thingID, randomParagraph)
		
		if err != nil {
			log.Fatalf("failed to edit comment: %v", err)
		}
	}

	
}

func generateRandomParagraph() string {
	f := faker.New()
	var paragraph string
    
	person := f.Person()
	paragraph += fmt.Sprintf("Hello, %s\n\n", person.Name())

	paragraph += fmt.Sprintf("You are %d years old\n\n", f.IntBetween(0, 100))

	phone := f.Phone()
	paragraph += fmt.Sprintf("Your phone number is %s\n\n", phone.Number())

	internet := f.Internet()
	paragraph += fmt.Sprintf("Your email is %s\n\n", internet.Email())

	company := f.Company()
	paragraph += fmt.Sprintf("You work for %s\n\n", company.Name())

	color := f.Color()
	paragraph += fmt.Sprintf("Your favorite color is %s\n\n", color.ColorName())

	payment := f.Payment()
	paragraph += fmt.Sprintf("Your credit card number is %s\n\n", payment.CreditCardNumber())

	friend := f.Person()
	lorem := f.Lorem()
	paragraph += fmt.Sprintf("Your friend %s sent you this message: '%s'\n\n", friend.Name(), lorem.Sentence(5))

	return paragraph
}
