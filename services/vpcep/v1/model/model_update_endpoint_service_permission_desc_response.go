package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEndpointServicePermissionDescResponse Response Object
type UpdateEndpointServicePermissionDescResponse struct {
	Permissions    *[]EpsPermission `json:"permissions,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o UpdateEndpointServicePermissionDescResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointServicePermissionDescResponse struct{}"
	}

	return strings.Join([]string{"UpdateEndpointServicePermissionDescResponse", string(data)}, " ")
}
