package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePermissionSetResponse Response Object
type CreatePermissionSetResponse struct {
	PermissionSet  *PermissionSetDto `json:"permission_set,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o CreatePermissionSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePermissionSetResponse struct{}"
	}

	return strings.Join([]string{"CreatePermissionSetResponse", string(data)}, " ")
}
