package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {

	// fetching all the records those sent status is not active
	filter := bson.D{{Key: "isSent", Value: false}}
	cursor, err := mongoClient.Database(APPDB).Collection("dashboard").Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	var results []bson.M
	if err = cursor.All(context.Background(), &results); err != nil {
		panic(err)
	}

	// iterated over all the records
	for _, result := range results {
		res, _ := json.Marshal(result)

		var mailer mailerModel
		err := json.Unmarshal(res, &mailer)

		if err != nil {
			log.Fatal(err)
		}

		// prepared the mail request data
		request := Mail{
			To:      mailer.ToAddress,
			Cc:      mailer.CcAddress,
			Subject: mailer.Subject,
			Body:    mailer.EmailBody,
		}
		// finally sent the email and showing the logs into the console
		sendEmail(request)
		fmt.Println(request)
		fmt.Println("Email sent to: " + strings.Join(mailer.ToAddress, ";"))
	}

	if err := mongoClient.Disconnect(context.TODO()); err != nil {
		log.Fatal("mongo database has disconnected")
	}
}
