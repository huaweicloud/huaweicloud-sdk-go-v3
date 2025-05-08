package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePostPaidKafkaInstanceRequest Request Object
type CreatePostPaidKafkaInstanceRequest struct {
	Body *CreateInstanceByEngineReq `json:"body,omitempty"`
}

func (o CreatePostPaidKafkaInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePostPaidKafkaInstanceRequest struct{}"
	}

	return strings.Join([]string{"CreatePostPaidKafkaInstanceRequest", string(data)}, " ")
}
