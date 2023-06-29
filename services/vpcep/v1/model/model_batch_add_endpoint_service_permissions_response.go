package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddEndpointServicePermissionsResponse Response Object
type BatchAddEndpointServicePermissionsResponse struct {
	Permissions    *[]EpsPermission `json:"permissions,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o BatchAddEndpointServicePermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddEndpointServicePermissionsResponse struct{}"
	}

	return strings.Join([]string{"BatchAddEndpointServicePermissionsResponse", string(data)}, " ")
}
