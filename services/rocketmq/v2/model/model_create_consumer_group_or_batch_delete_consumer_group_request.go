package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConsumerGroupOrBatchDeleteConsumerGroupRequest Request Object
type CreateConsumerGroupOrBatchDeleteConsumerGroupRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 批量删除消费组时使用，不配置则为创建接口。删除操作：delete。
	Action *string `json:"action,omitempty"`

	Body *CreateConsumerGroupOrBatchDeleteConsumerGroupReq `json:"body,omitempty"`
}

func (o CreateConsumerGroupOrBatchDeleteConsumerGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConsumerGroupOrBatchDeleteConsumerGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateConsumerGroupOrBatchDeleteConsumerGroupRequest", string(data)}, " ")
}
