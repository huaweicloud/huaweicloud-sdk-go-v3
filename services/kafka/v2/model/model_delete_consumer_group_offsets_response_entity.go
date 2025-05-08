package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteConsumerGroupOffsetsResponseEntity struct {

	// topic名称
	Name string `json:"name"`

	// 消费位点删除是否成功
	Success bool `json:"success"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`
}

func (o DeleteConsumerGroupOffsetsResponseEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConsumerGroupOffsetsResponseEntity struct{}"
	}

	return strings.Join([]string{"DeleteConsumerGroupOffsetsResponseEntity", string(data)}, " ")
}
