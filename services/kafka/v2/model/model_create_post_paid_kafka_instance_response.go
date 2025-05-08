package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePostPaidKafkaInstanceResponse Response Object
type CreatePostPaidKafkaInstanceResponse struct {

	// 实例ID
	InstanceId     *string `json:"instance_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePostPaidKafkaInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePostPaidKafkaInstanceResponse struct{}"
	}

	return strings.Join([]string{"CreatePostPaidKafkaInstanceResponse", string(data)}, " ")
}
