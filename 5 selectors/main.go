// Here i am using 5 selectors they are $eq, $gt, $regex, $ne, $ in
// here i will explain where i used the selectors
// $eq: here i used for to find the documents where a the order documents are stored
// $gt: i used for find documents where a field is greater than a value for the below code i gave 1000
// $regex: i used check the regular expression of a apple order for the below code
// $ne:  here i used the command  for see the cancellation of a order
// $in : for example there are products collection in there are some attributes from that attributes selecting a documents for example from the below code

// Here i am taking the order data from mongo DB by using 5 selectors as mentioned

package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// connection to the MongoDB database
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") // local host database that documents containing
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// handle to the "orders" collection
	collection := client.Database("mydb").Collection("orders")

	// Find documents whether the "status" field equals "delivered"
	filter1 := bson.M{"status": bson.M{"$eq": "delivered"}}
	cursor1, err := collection.Find(context.Background(), filter1)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor1.Close(context.Background())

	// Find documents where the "total" field is greater than 1000
	filter2 := bson.M{"total": bson.M{"$gt": 1000}}
	cursor2, err := collection.Find(context.Background(), filter2)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor2.Close(context.Background())

	// Find documents where the "products.name" field contains "apple"
	filter3 := bson.M{"products.name": bson.M{"$regex": "apple"}}
	cursor3, err := collection.Find(context.Background(), filter3)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor3.Close(context.Background())

	// Find documents where the "customer.email" field ends with "@example.com"
	filter4 := bson.M{"customer.email": bson.M{"$regex": "vissarapusrinath@gmail.com$"}}
	cursor4, err := collection.Find(context.Background(), filter4)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor4.Close(context.Background())

	// Find documents where the "status" field is not equal to "cancelled"
	filter5 := bson.M{"status": bson.M{"$ne": "cancelled"}}
	cursor5, err := collection.Find(context.Background(), filter5)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor5.Close(context.Background())

	// Iterate over the cursors and print the documents
	for cursor1.Next(context.Background()) {
		var doc bson.M
		if err := cursor1.Decode(&doc); err != nil {
			log.Fatal(err)
		}
		fmt.Println(doc)
	}

	for cursor2.Next(context.Background()) {
		var doc bson.M
		if err := cursor2.Decode(&doc); err != nil {
			log.Fatal(err)
		}
		fmt.Println(doc)
	}

	for cursor3.Next(context.Background()) {
		var doc bson.M
		if err := cursor3.Decode(&doc); err != nil {
			log.Fatal(err)
		}
		fmt.Println(doc)
	}

	for cursor4.Next(context.Background()) {
		var doc bson.M
		if err := cursor4.Decode(&doc); err != nil {
			log.Fatal(err)
		}
		fmt.Println(doc)
	}

	for cursor5.Next(context.Background()) {
		var doc bson.M
		if err := cursor5.Decode(&doc); err != nil {
			log.Fatal(err)
		}
		fmt.Println(doc)
	}
}
