package controllers

import (
    "crypto/md5"
    "encoding/hex"
    "time"

    "fmt"
    "github.com/Kamva/mgm/v2"
    "github.com/gofiber/fiber"
    "github.com/parvusvox/poseServer_v2/models"
    "go.mongodb.org/mongo-driver/bson"
    
    "os"
    "strconv"
)

func GetFrames(ctx *fiber.Ctx){
    passphrase := os.Getenv("PASSPHRASE")

    token := ctx.Params("token")
    recId := ctx.Params("recId")

    location,err := time.LoadLocation("UTC")
    t := time.Now().In(location)
    hr, min, _ := t.Clock()
    toHash := "" + strconv.Itoa(int(t.Month())) + strconv.Itoa(int(t.Day())) + strconv.Itoa(int(hr)) + strconv.Itoa(int(min))
    toHash = recId + toHash + passphrase

    var hash = make(chan []byte, 1)
    sum := md5.Sum([]byte(toHash))
    hash <- sum[:]

    hashString := hex.EncodeToString(<-hash)

    fmt.Println(hashString)

    if(true || token == hashString){
        collection := mgm.Coll(&models.Frame{})
        frames := []models.Frame{}

        err = collection.SimpleFind(&frames, bson.M{"recId": recId})
        if err != nil{
            ctx.Status(500).JSON(fiber.Map{
                "ok": false,
                "error": err.Error(),
            })
            return
        }
        ctx.JSON(frames)
    } else{
        ctx.JSON(fiber.Map{
            "ok": false,
            "error": "token mismatch",
        })
    }
}


func CreateFrame(ctx *fiber.Ctx){
    passphrase := os.Getenv("PASSPHRASE")

    params := new(struct{
        Token string
        RecId string
        Frame struct {
            Filename string
            Pose string
        }
    })

    ctx.BodyParser(&params)

    tohash := passphrase + params.Frame.Filename + params.Frame.Pose

    var hash = make(chan []byte, 1)
    sum := md5.Sum([]byte(tohash))
    hash <- sum[:]

    hashString := hex.EncodeToString(<-hash)

    fmt.Println(hashString)

    if params.Token == hashString {
        frame := models.CreateFrame(params.Token,
            params.RecId,
            params.Frame.Filename,
            params.Frame.Pose)

        err := mgm.Coll(frame).Create(frame)

        if err != nil{
            ctx.Status(500).JSON(fiber.Map{
                "ok": false,
                "error": err.Error(),
            })
            return
        }

        ctx.JSON(fiber.Map{
            "ok": true,
        })
    } else{
        ctx.JSON(fiber.Map{
            "ok": false,
            "error": "token mismatch",
        })
    }
}
