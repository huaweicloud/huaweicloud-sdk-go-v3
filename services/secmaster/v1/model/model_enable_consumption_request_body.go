package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EnableConsumptionRequestBody struct {

	// 数据管道ID
	PipeId string `json:"pipe_id"`
}

func (o EnableConsumptionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableConsumptionRequestBody struct{}"
	}

	return strings.Join([]string{"EnableConsumptionRequestBody", string(data)}, " ")
}
