package main

import (
    "github.com/vartanbeno/go-reddit/v2/reddit"
    "log"
	"context"
	"fmt"
	"github.com/jaswdr/faker"
	"time"
	"encoding/json"
	"os"
)

type Secret struct {
	ID       string `json:"ID"`
	Secret   string `json:"Secret"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}

func main() {
	var numberofComments int

    fmt.Println("Enter number of comments to delete: ")
    _, err := fmt.Scanf("%d", &numberofComments )
	
	file, err := os.Open("secret.json")
	if err != nil{
      log.Fatalf("fail to open file: %v", err)
	}

	defer file.Close()
    
	decoder := json.NewDecoder(file)
	secrets := Secret{}
    err = decoder.Decode(&secrets)

	if err != nil {
		log.Fatalf("Error decoding config file: %v", err)
	}

	// fmt.Println(secrets)

	credentials := reddit.Credentials{
		ID:       secrets.ID,
		Secret:   secrets.Secret,
		Username: secrets.Username,
		Password: secrets.Password,
	}

    client, err := reddit.NewClient(credentials)
    if err != nil {
        log.Fatalf("failed to create client: %v", err)
    }

	// Assuming 'client' is already authenticated
	comments, _, err := client.User.Comments(context.Background(), &reddit.ListUserOverviewOptions{
		ListOptions: reddit.ListOptions{
			Limit: numberofComments,
		},
	})
	if err != nil {
		log.Fatalf("failed to fetch comments: %v", err)
	}
	fmt.Println(comments)
	randomParagraph := fmt.Sprintf(generateRandomParagraph())

	for _, comment := range comments {
		fmt.Println(comment.Body)
		thingID := "t1_" + comment.ID
		_, _, err := client.Comment.Edit(context.Background(), thingID, randomParagraph)

		if err != nil {
			log.Fatalf("failed to edit comment: %v", err)
		} else {
			fmt.Println("edit successful")
		}

		time.Sleep(10 * time.Second)

		_, err = client.Comment.Delete(context.Background(), thingID)

		if err != nil{
			log.Fatalf("failed to delete comment: %v", err)
		} else {
			fmt.Println("delete successful")
		}

		time.Sleep(10 * time.Second)

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
