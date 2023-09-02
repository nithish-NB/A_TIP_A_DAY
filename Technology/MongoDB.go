package tech

import (
	data "demo/Data"
	"fmt"
	"log"
)

func Mongo1() {
	fmt.Println("MongoDB Technical Tips with Source Links:")
	client, ctx := Context()
	database := client.Database("Tools")
	Mongodata := database.Collection("MongoDB")
	//	var Javatool data.Tip
	var abc []data.Tip
	// Tips stored in a slice
	mongodbTips := []string{
		"1. Start with the official MongoDB documentation to understand the basics. [MongoDB Documentation](https://docs.mongodb.com/)",
		"2. Install and set up MongoDB following the official installation guide. [Installation Guide](https://docs.mongodb.com/manual/installation/)",
		"3. Use a MongoDB GUI client like MongoDB Compass for visual management. [MongoDB Compass](https://www.mongodb.com/try/download/compass)",
		"4. Create databases and collections to organize your data effectively. [Create Databases and Collections](https://docs.mongodb.com/manual/core/databases-and-collections/)",
		"5. Understand BSON, MongoDB's binary JSON format, for data storage. [BSON](https://docs.mongodb.com/manual/reference/bson-types/)",
		"6. Master CRUD operations: Create, Read, Update, and Delete documents. [CRUD Operations](https://docs.mongodb.com/manual/crud/)",
		"7. Utilize indexes to improve query performance. [Indexes](https://docs.mongodb.com/manual/indexes/)",
		"8. Learn aggregation pipelines for complex data transformations. [Aggregation](https://docs.mongodb.com/manual/aggregation/)",
		"9. Practice schema design based on your application's needs. [Data Modeling](https://docs.mongodb.com/manual/core/data-modeling-introduction/)",
		"10. Understand the benefits and use cases of document-oriented databases like MongoDB. [Document-Oriented Databases](https://docs.mongodb.com/manual/core/databases-and-collections/#document-oriented)",
		"11. Implement data validation using JSON Schema. [Data Validation](https://docs.mongodb.com/manual/core/schema-validation/)",
		"12. Secure your MongoDB instance with authentication and authorization. [Security](https://docs.mongodb.com/manual/security/)",
		"13. Set up user roles and access control to restrict permissions. [Role-Based Access Control](https://docs.mongodb.com/manual/core/security-built-in-roles/)",
		"14. Perform data backups and implement a disaster recovery plan. [Back Up and Restore](https://docs.mongodb.com/manual/core/backups/)",
		"15. Monitor MongoDB performance using tools like MongoDB Atlas or Ops Manager. [Monitoring](https://docs.mongodb.com/manual/administration/monitoring/)",
		"16. Tune MongoDB for optimal performance with configuration options. [Configuration](https://docs.mongodb.com/manual/reference/configuration-options/)",
		"17. Handle errors and exceptions effectively in your MongoDB application. [Error Handling](https://docs.mongodb.com/manual/core/error-messages/)",
		"18. Keep up-to-date with MongoDB releases and updates. [Release Notes](https://docs.mongodb.com/manual/release-notes/)",
		"19. Utilize MongoDB drivers and libraries for various programming languages. [Drivers](https://docs.mongodb.com/ecosystem/drivers/)",
		"20. Use sharding to distribute data across multiple servers for scalability. [Sharding](https://docs.mongodb.com/manual/sharding/)",
		"21. Explore MongoDB Atlas, the managed database service by MongoDB. [MongoDB Atlas](https://www.mongodb.com/cloud/atlas)",
		"22. Implement geospatial queries for location-based applications. [Geospatial Queries](https://docs.mongodb.com/manual/geospatial-queries/)",
		"23. Learn about change streams for real-time data updates. [Change Streams](https://docs.mongodb.com/manual/changeStreams/)",
		"24. Apply time-to-live (TTL) indexes for automatic document expiration. [TTL Indexes](https://docs.mongodb.com/manual/core/index-ttl/)",
		"25. Understand how MongoDB handles transactions. [Transactions](https://docs.mongodb.com/manual/core/transactions/)",
		"26. Optimize query performance using query hints. [Query Hints](https://docs.mongodb.com/manual/reference/query-hint/)",
		"27. Monitor and optimize your MongoDB deployment for storage efficiency. [Storage](https://docs.mongodb.com/manual/core/storage/)",
		"28. Utilize the MongoDB Atlas Data Lake for data lake capabilities. [Data Lake](https://docs.mongodb.com/datalake/)",
		"29. Stay informed about MongoDB best practices for development and operations. [Best Practices](https://docs.mongodb.com/manual/best-practices/)",
		"30. Engage with the MongoDB community and forums for support and knowledge sharing. [Community](https://www.mongodb.com/community)",
	}

	// Print the MongoDB tips
	for i, tip := range mongodbTips {
		abc = append(abc, data.Tip{Date: i + 1, Technology: "MongoDb", Content: tip})
		userToInsert := abc[i]
		_, err := Mongodata.InsertOne(ctx, userToInsert)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Registered  successfully")
	}
}
