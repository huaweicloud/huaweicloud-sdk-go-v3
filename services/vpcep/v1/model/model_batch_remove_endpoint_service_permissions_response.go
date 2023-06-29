package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRemoveEndpointServicePermissionsResponse Response Object
type BatchRemoveEndpointServicePermissionsResponse struct {
	Permissions    *[]EpsPermission `json:"permissions,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o BatchRemoveEndpointServicePermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRemoveEndpointServicePermissionsResponse struct{}"
	}

	return strings.Join([]string{"BatchRemoveEndpointServicePermissionsResponse", string(data)}, " ")
}
