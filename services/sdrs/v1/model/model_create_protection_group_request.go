package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateProtectionGroupRequest struct {
	Body *CreateProtectionGroupRequestBody `json:"body,omitempty"`
}

func (o CreateProtectionGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProtectionGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateProtectionGroupRequest", string(data)}, " ")
}
