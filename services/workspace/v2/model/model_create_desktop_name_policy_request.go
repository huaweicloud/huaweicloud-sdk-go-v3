package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDesktopNamePolicyRequest Request Object
type CreateDesktopNamePolicyRequest struct {
	Body *CreateDesktopNamePolicyReq `json:"body,omitempty"`
}

func (o CreateDesktopNamePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDesktopNamePolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateDesktopNamePolicyRequest", string(data)}, " ")
}
