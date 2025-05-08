package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteConsumerGroupOffsetsResponse Response Object
type DeleteConsumerGroupOffsetsResponse struct {

	// 结果列表
	Topics         *[]DeleteConsumerGroupOffsetsResponseEntity `json:"topics,omitempty"`
	HttpStatusCode int                                         `json:"-"`
}

func (o DeleteConsumerGroupOffsetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConsumerGroupOffsetsResponse struct{}"
	}

	return strings.Join([]string{"DeleteConsumerGroupOffsetsResponse", string(data)}, " ")
}
