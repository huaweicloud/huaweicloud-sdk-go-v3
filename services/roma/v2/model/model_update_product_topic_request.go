package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateProductTopicRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 产品ID

	ProductId int32 `json:"product_id"`
	// 产品主题ID

	TopicId int32 `json:"topic_id"`

	Body *UpdateProductTopicRequestBody `json:"body,omitempty"`
}

func (o UpdateProductTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProductTopicRequest struct{}"
	}

	return strings.Join([]string{"UpdateProductTopicRequest", string(data)}, " ")
}
