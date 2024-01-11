package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecurityPermissionSetRequest Request Object
type CreateSecurityPermissionSetRequest struct {

	// DataArts Studio工作空间ID
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
