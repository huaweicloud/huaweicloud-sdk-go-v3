package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTuningResultRequestBody 获取调优结果请求体
type ShowTuningResultRequestBody struct {

	// SQL诊断消息ID
	MessageId string `json:"message_id"`
}

func (o ShowTuningResultRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTuningResultRequestBody struct{}"
	}

	return strings.Join([]string{"ShowTuningResultRequestBody", string(data)}, " ")
}
