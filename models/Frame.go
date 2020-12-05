package models

import (
    "github.com/Kamva/mgm/v2"
)

type FrameData struct{
    Filename string `json:"filename" bson:"filename"`
    Pose string `json:"pose" bson:"pose"`
}

type Frame struct{
    mgm.DefaultModel `bson:",inline"`
    Token string `json:"title" bson:"title"`
    RecId string `json:"recId" bson:"recId"`
    Frame FrameData `json:"frame" bson:"frame"`
}

func CreateFrame(token, recId, filename, pose string) *Frame{
    return &Frame{
        Token : token,
        RecId : recId,
        Frame : FrameData {
            Filename: filename,
            Pose: pose,
        },
    }
}
