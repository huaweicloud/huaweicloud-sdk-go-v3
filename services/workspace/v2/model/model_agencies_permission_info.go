package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgenciesPermissionInfo 委托权限信息
type AgenciesPermissionInfo struct {

	// 委托权限项
	SystemPermissionDisplayNames *[]string `json:"system_permission_display_names,omitempty"`

	// 需要委托的权限项
	WantedSystemPermissionDisplayNames *[]string `json:"wanted_system_permission_display_names,omitempty"`
}

func (o AgenciesPermissionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgenciesPermissionInfo struct{}"
	}

	return strings.Join([]string{"AgenciesPermissionInfo", string(data)}, " ")
}
