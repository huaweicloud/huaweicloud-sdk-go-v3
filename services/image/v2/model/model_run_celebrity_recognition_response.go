package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunCelebrityRecognitionResponse Response Object
type RunCelebrityRecognitionResponse struct {

	// 调用成功时表示调用结果。  调用失败时无此字段。
	Result         *[]CelebrityRecognitionResultBody `json:"result,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o RunCelebrityRecognitionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunCelebrityRecognitionResponse struct{}"
	}

	return strings.Join([]string{"RunCelebrityRecognitionResponse", string(data)}, " ")
}
