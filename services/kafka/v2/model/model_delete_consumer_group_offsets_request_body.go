package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteConsumerGroupOffsetsRequestBody struct {

	// topic列表
	Topics *[]string `json:"topics,omitempty"`
}

func (o DeleteConsumerGroupOffsetsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConsumerGroupOffsetsRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteConsumerGroupOffsetsRequestBody", string(data)}, " ")
}
