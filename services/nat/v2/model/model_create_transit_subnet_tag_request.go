package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTransitSubnetTagRequest Request Object
type CreateTransitSubnetTagRequest struct {

	// 中转子网的ID。
	ResourceId string `json:"resource_id"`

	Body *CreateResourceTagRequestBody `json:"body,omitempty"`
}

func (o CreateTransitSubnetTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTransitSubnetTagRequest struct{}"
	}

	return strings.Join([]string{"CreateTransitSubnetTagRequest", string(data)}, " ")
}
