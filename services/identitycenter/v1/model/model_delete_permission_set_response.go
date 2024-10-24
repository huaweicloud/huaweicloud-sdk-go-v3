package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePermissionSetResponse Response Object
type DeletePermissionSetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePermissionSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePermissionSetResponse struct{}"
	}

	return strings.Join([]string{"DeletePermissionSetResponse", string(data)}, " ")
}
