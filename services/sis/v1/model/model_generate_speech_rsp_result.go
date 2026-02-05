package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GenerateSpeechRspResult 调用成功表示识别结果，调用失败时无此字段。
type GenerateSpeechRspResult struct {

	// 语音数据，Base64编码格式返回。
	Data *string `json:"data,omitempty"`
}

func (o GenerateSpeechRspResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GenerateSpeechRspResult struct{}"
	}

	return strings.Join([]string{"GenerateSpeechRspResult", string(data)}, " ")
}
