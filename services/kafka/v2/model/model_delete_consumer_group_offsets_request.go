package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteConsumerGroupOffsetsRequest Request Object
type DeleteConsumerGroupOffsetsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 消费组ID。
	Group string `json:"group"`

	Body *DeleteConsumerGroupOffsetsRequestBody `json:"body,omitempty"`
}

func (o DeleteConsumerGroupOffsetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConsumerGroupOffsetsRequest struct{}"
	}

	return strings.Join([]string{"DeleteConsumerGroupOffsetsRequest", string(data)}, " ")
}
