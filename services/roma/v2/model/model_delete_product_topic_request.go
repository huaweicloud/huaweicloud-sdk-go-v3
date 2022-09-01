package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteProductTopicRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 产品ID
	ProductId int32 `json:"product_id" xml:"product_id"`

	// 产品主题ID
	TopicId int32 `json:"topic_id" xml:"topic_id"`
}

func (o DeleteProductTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProductTopicRequest struct{}"
	}

	return strings.Join([]string{"DeleteProductTopicRequest", string(data)}, " ")
}
