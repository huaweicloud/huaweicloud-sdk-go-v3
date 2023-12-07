package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecurityPermissionSetRequest Request Object
type CreateSecurityPermissionSetRequest struct {

	// workspace 信息
	Workspace string `json:"workspace"`

	Body *PermissionSetCreateDto `json:"body,omitempty"`
}

func (o CreateSecurityPermissionSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityPermissionSetRequest struct{}"
	}

	return strings.Join([]string{"CreateSecurityPermissionSetRequest", string(data)}, " ")
}
