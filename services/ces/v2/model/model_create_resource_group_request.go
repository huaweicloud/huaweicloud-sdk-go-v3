package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResourceGroupRequest Request Object
type CreateResourceGroupRequest struct {
	Body *CreateResourceGroupRequestBody `json:"body,omitempty"`
}

func (o CreateResourceGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateResourceGroupRequest", string(data)}, " ")
}
