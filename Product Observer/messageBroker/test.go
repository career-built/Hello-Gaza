package messageBroker

import (
	// "example/baseProject/api"
	// "example/baseProject/database"
	// "example/baseProject/product"

	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	// "log"
	// "github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Starting Base")
	// dbConnector := database.NewPostgres()
	// if dbConnector == nil {
	// 	log.Fatal("can't connect to database")
	// }
	// defer dbConnector.CloseDB()

	// //Path the Database to the base Db interface
	// productMgr := product.NewProductService(dbConnector)

	// //Path the product feature to the base manger interface
	// productRouter := api.NewProductRouter(productMgr)

	// e := echo.New()
	// e.POST("/product/create", productRouter.CreateProduct)
	// e.GET("/product/:id", productRouter.GetProductByID)

	// e.Logger.Fatal(e.Start(":3030"))
	broker, err := NewRabbitMQBroker("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer broker.Close()

	queueName := "messages"

	// // Publish messages
	// messages := []string{"Hello, RabbitMQ!", "Another message"}
	// err = broker.PublishMessages("", queueName, messages)
	// if err != nil {
	// 	panic(err)
	// }

	// Consume messages
	err = broker.ConsumeMessages(queueName, func(message string) {
		fmt.Println("Received a message:", message)
	})
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter a message (or 'exit' to quit): ")
		scanner.Scan()
		message := scanner.Text()

		if strings.ToLower(message) == "exit" {
			break
		}
		err = broker.PublishMessages("", queueName, []string{message})
		if err != nil {
			fmt.Println("Message NOT sent:", message)
			log.Fatalf("%s: %s", message, err)
		}
		fmt.Println("Message sent:", message)
	}
	// Keep the program running (you may need to stop it manually)
	// select {}
}
