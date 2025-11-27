package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDistributionRequest Request Object
type CreateDistributionRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *CreateDistributionRequestBody `json:"body,omitempty"`
}

func (o CreateDistributionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDistributionRequest struct{}"
	}

	return strings.Join([]string{"CreateDistributionRequest", string(data)}, " ")
}
