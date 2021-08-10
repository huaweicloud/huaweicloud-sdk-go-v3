package model

import (
	"encoding/json"

	"strings"
)

//
type CelebrityRecognitionResultBody struct {
	// 置信度，取值范围 0-1。

	Confidence *float32 `json:"confidence,omitempty"`
	// 名人的面部信息，包括4个值：  h：人脸区域高度  w：人脸区域宽度  x：人脸区域左上角到y轴距离  y：人脸区域左上角到x轴距离

	FaceDetail *interface{} `json:"face_detail,omitempty"`
	// label为对应的名人信息。

	Label *string `json:"label,omitempty"`
}

func (o CelebrityRecognitionResultBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CelebrityRecognitionResultBody struct{}"
	}

	return strings.Join([]string{"CelebrityRecognitionResultBody", string(data)}, " ")
}
