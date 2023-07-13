package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePermissionSetResponse Response Object
type UpdatePermissionSetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdatePermissionSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePermissionSetResponse struct{}"
	}

	return strings.Join([]string{"UpdatePermissionSetResponse", string(data)}, " ")
}
