package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateApplicationGroupRequest Request Object
type CreateApplicationGroupRequest struct {
	Body *GroupCreateRequest `json:"body,omitempty"`
}

func (o CreateApplicationGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApplicationGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateApplicationGroupRequest", string(data)}, " ")
}
