package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddEndpointPermissionsResponse Response Object
type AddEndpointPermissionsResponse struct {

	// 权限列表
	Permissions *[]string `json:"permissions,omitempty"`

	XRequestId     *string `json:"x-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddEndpointPermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddEndpointPermissionsResponse struct{}"
	}

	return strings.Join([]string{"AddEndpointPermissionsResponse", string(data)}, " ")
}
