package tech

import (
	"context"
	data "demo/Data"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Context() (*mongo.Client, context.Context) {
	ctx := context.TODO()
	clientoptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientoptions)
	if err != nil {
		log.Fatal(err)
	}
	return client, ctx
}
func Java1() {
	fmt.Println("Java Technical Tips with Source Links:")
	client, ctx := Context()
	database := client.Database("Tools")
	java := database.Collection("Java")
	//	var Javatool data.Tip
	var abc []data.Tip
	defer client.Disconnect(ctx)
	// Tips stored in a slice
	javaTips := []string{
		"1. Start with the official Java documentation to understand the basics. [Java Documentation](https://docs.oracle.com/javase/)",
		"2. Use a development environment like Eclipse, IntelliJ IDEA, or Visual Studio Code for coding. [Eclipse](https://www.eclipse.org/), [IntelliJ IDEA](https://www.jetbrains.com/idea/), [Visual Studio Code](https://code.visualstudio.com/)",
		"3. Master object-oriented programming (OOP) concepts like classes, objects, and inheritance. [Object-Oriented Programming](https://docs.oracle.com/javase/tutorial/java/concepts/)",
		"4. Practice coding exercises on platforms like LeetCode and HackerRank. [LeetCode](https://leetcode.com/), [HackerRank](https://www.hackerrank.com/domains/tutorials/10-days-of-java)",
		"5. Understand the importance of memory management and garbage collection in Java. [Memory Management](https://docs.oracle.com/javase/guides/troubleshoot/gctuning/)",
		"6. Learn about data structures like arrays, lists, and maps for efficient data manipulation. [Data Structures](https://docs.oracle.com/javase/tutorial/collections/intro/)",
		"7. Use exception handling to gracefully handle errors and exceptions. [Exception Handling](https://docs.oracle.com/javase/tutorial/essential/exceptions/)",
		"8. Familiarize yourself with Java libraries and frameworks for various application domains. [Java Libraries and Frameworks](https://www.baeldung.com/java-libraries)",
		"9. Practice multi-threading and concurrency for parallel processing. [Concurrency](https://docs.oracle.com/javase/tutorial/essential/concurrency/)",
		"10. Implement design patterns for solving common software design problems. [Design Patterns](https://www.tutorialspoint.com/design_pattern/index.htm)",
		"11. Optimize your Java code using profiling tools like VisualVM. [VisualVM](https://visualvm.github.io/)",
		"12. Use version control systems like Git for code management. [Git](https://git-scm.com/)",
		"13. Write unit tests using JUnit or TestNG to ensure code correctness. [JUnit](https://junit.org/), [TestNG](https://testng.org/)",
		"14. Leverage the Java Streams API for data processing and manipulation. [Streams API](https://docs.oracle.com/javase/8/docs/api/java/util/stream/package-summary.html)",
		"15. Explore JavaFX for building desktop applications with a modern user interface. [JavaFX](https://openjfx.io/)",
		"16. Learn about Java Database Connectivity (JDBC) for database interactions. [JDBC](https://docs.oracle.com/javase/tutorial/jdbc/)",
		"17. Stay updated with Java releases and new features. [Java Release Notes](https://www.oracle.com/java/technologies/javase-downloads.html)",
		"18. Practice code refactoring to improve code quality and maintainability. [Code Refactoring](https://refactoring.com/)",
		"19. Understand dependency injection and inversion of control (IoC) for modular and testable code. [Dependency Injection](https://docs.oracle.com/javaee/5/tutorial/doc/bnbpz.html)",
		"20. Implement effective logging using libraries like Log4j or SLF4J. [Log4j](https://logging.apache.org/log4j/), [SLF4J](http://www.slf4j.org/)",
		"21. Use the Java NIO package for non-blocking I/O operations. [Java NIO](https://docs.oracle.com/javase/tutorial/essential/io/fileio.html)",
		"22. Familiarize yourself with software development methodologies like Agile and Scrum. [Agile](https://www.agilealliance.org/agile101/), [Scrum](https://www.scrum.org/resources/what-is-scrum)",
		"23. Practice clean code principles for readability and maintainability. [Clean Code](https://cleancoders.com/)",
		"24. Learn about internationalization and localization for global applications. [Internationalization](https://docs.oracle.com/javase/tutorial/i18n/index.html)",
		"25. Keep your skills up-to-date with Java-related blogs and online courses. [Java Blogs](https://dzone.com/java-jdk-development-tutorials-tools-news), [Online Courses](https://www.coursera.org/courses?query=java)",
		"26. Master the use of lambdas and functional programming features introduced in Java 8. [Java 8 Features](https://docs.oracle.com/javase/8/docs/technotes/guides/language/index.html)",
		"27. Implement RESTful web services using Java technologies like Spring Boot. [Spring Boot](https://spring.io/projects/spring-boot)",
		"28. Secure your Java applications using best practices and security libraries. [Java Security](https://docs.oracle.com/javase/tutorial/security/)",
		"29. Practice test-driven development (TDD) to ensure code reliability. [Test-Driven Development](https://www.agilealliance.org/glossary/tdd/)",
		"30. Engage with the Java community through forums, conferences, and user groups. [Java Community](https://community.oracle.com/community/developer)",
	}

	// Print the Java tips
	for i, tip := range javaTips {
		abc = append(abc, data.Tip{Date: i + 1, Technology: "Java", Content: tip})
		userToInsert := abc[i]
		_, err := java.InsertOne(ctx, userToInsert)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Registered  successfully")
	}
}
