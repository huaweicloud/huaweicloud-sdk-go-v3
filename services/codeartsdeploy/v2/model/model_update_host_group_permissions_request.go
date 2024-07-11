package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHostGroupPermissionsRequest Request Object
type UpdateHostGroupPermissionsRequest struct {

	// 主机集群id
	GroupId string `json:"group_id"`

	Body *PermissionUpdateV2Body `json:"body,omitempty"`
}

func (o UpdateHostGroupPermissionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHostGroupPermissionsRequest struct{}"
	}

	return strings.Join([]string{"UpdateHostGroupPermissionsRequest", string(data)}, " ")
}
