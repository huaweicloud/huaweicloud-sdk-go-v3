package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEndpointPermissionsResponse Response Object
type DeleteEndpointPermissionsResponse struct {

	// 权限列表
	Permissions *[]string `json:"permissions,omitempty"`

	XRequestId     *string `json:"x-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteEndpointPermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEndpointPermissionsResponse struct{}"
	}

	return strings.Join([]string{"DeleteEndpointPermissionsResponse", string(data)}, " ")
}
