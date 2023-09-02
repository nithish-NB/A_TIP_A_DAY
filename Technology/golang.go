package tech

import (
	data "demo/Data"
	"fmt"
	"log"
)

func GO1() {
	//fmt.Println("Go (Golang) Tips for Beginners:")
	client, ctx := Context()
	database := client.Database("Tools")
	godata := database.Collection("Go")
	//	var Javatool data.Tip
	var abc []data.Tip
	defer client.Disconnect(ctx)
	// Tips stored in a slice
	// Tips stored in a slice
	golangTips := []string{
		"1. Start with the official Go documentation to understand the basics. [Go Documentation](https://golang.org/doc/)",
		"2. Use 'go fmt' to automatically format your code according to Go's style guide. [gofmt](https://golang.org/cmd/gofmt/)",
		"3. Keep your code clean and simple; Go promotes readability and simplicity.",
		"4. Learn the different data types in Go, such as int, string, and bool. [Basic Types](https://tour.golang.org/basics/11)",
		"5. Master variable declaration and assignment in Go. [Variables](https://tour.golang.org/basics/8)",
		"6. Practice control structures like if, else, for, and switch. [Flow Control](https://tour.golang.org/flowcontrol/1)",
		"7. Use named return values for functions to improve code clarity. [Named Results](https://tour.golang.org/basics/7)",
		"8. Understand slices and arrays for working with collections of data. [Slices](https://tour.golang.org/moretypes/7)",
		"9. Explore goroutines and channels for concurrent programming. [Concurrency](https://tour.golang.org/concurrency/1)",
		"10. Learn about defer and how it's used for cleanup operations. [Defer](https://tour.golang.org/flowcontrol/12)",
		"11. Handle errors explicitly using the 'error' type and 'if err != nil' checks. [Errors](https://tour.golang.org/methods/9)",
		"12. Familiarize yourself with Go's package system for code organization. [Packages](https://tour.golang.org/basics/1)",
		"13. Use 'go get' to fetch external packages and 'go mod' for dependency management. [Packages and Modules](https://golang.org/doc/code.html)",
		"14. Write unit tests for your functions using the 'testing' package. [Testing](https://golang.org/pkg/testing/)",
		"15. Leverage the 'fmt' package for input and output. [fmt](https://pkg.go.dev/fmt)",
		"16. Understand pointers and their use cases. [Pointers](https://tour.golang.org/moretypes/1)",
		"17. Explore struct types for creating custom data structures. [Structs](https://tour.golang.org/moretypes/2)",
		"18. Implement interfaces to define behavior shared by multiple types. [Interfaces](https://tour.golang.org/methods/1)",
		"19. Use maps for key-value data storage. [Maps](https://tour.golang.org/moretypes/9)",
		"20. Study error handling patterns like 'panic' and 'recover'. [Panic and Recover](https://blog.golang.org/defer-panic-and-recover)",
		"21. Experiment with closures and anonymous functions. [Function Closures](https://tour.golang.org/moretypes/25)",
		"22. Use the 'range' keyword for iterating over slices, maps, and channels. [For Range](https://tour.golang.org/moretypes/16)",
		"23. Write clear and concise function and variable names.",
		"24. Avoid global variables whenever possible.",
		"25. Follow the naming conventions for exported and unexported (private) names.",
		"26. Learn from the Go community by reading open-source Go code.",
		"27. Utilize code linters like 'golint' or 'golangci-lint' to catch issues early.",
		"28. Keep up with Go's evolving ecosystem and best practices.",
		"29. Embrace the Go philosophy: 'Do not communicate by sharing memory; instead, share memory by communicating.' [Share Memory By Communicating](https://blog.golang.org/share-memory-by-communicating)",
		"30. Practice consistently; regular coding exercises and projects will help you improve.",
	}

	// Print the Go tips
	for i, tip := range golangTips {
		abc = append(abc, data.Tip{Date: i + 1, Technology: "GO", Content: tip})
		userToInsert := abc[i]
		_, err := godata.InsertOne(ctx, userToInsert)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Registered  successfully")
	}
}
