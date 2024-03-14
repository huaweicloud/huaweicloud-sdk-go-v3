package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEndpointServicePermissionDescResponse Response Object
type UpdateEndpointServicePermissionDescResponse struct {

	// 终端节点服务白名单
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
