package main

import(
    "log"
    "fmt"
    "os"

    "github.com/Kamva/mgm/v2"
    "github.com/gofiber/fiber"
    "github.com/parvusvox/poseServer_v2/controllers"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func init(){
    connectionString := os.Getenv("CONNSTRING")
    err := mgm.SetDefaultConfig(nil, "data", options.Client().ApplyURI(connectionString))
    if err != nil{
        log.Fatal(err)
    }
}

func main(){
    app := fiber.New()

    fmt.Println("SERVER UP")
    app.Get("/", func(ctx *fiber.Ctx){
        ctx.Send("COGNALITY POSE SERVER API Version 2.0.0\nrewritten in beautiful golang")
    })

    app.Get("/frames/:token/:recId", controllers.GetFrames)
    app.Post("/frames", controllers.CreateFrame)

    port := os.Getenv("PORT")
    if port == ""{
        port = "3000"
    }

    app.Listen(port)
}
