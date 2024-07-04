package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBindingsRequest Request Object
type ListBindingsRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// Vhost名称
	Vhost string `json:"vhost"`

	// Exchange名称
	Exchange string `json:"exchange"`
}

func (o ListBindingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBindingsRequest struct{}"
	}

	return strings.Join([]string{"ListBindingsRequest", string(data)}, " ")
}
