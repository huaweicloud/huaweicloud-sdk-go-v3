package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgenciesInfo 委托信息
type AgenciesInfo struct {

	// 名称
	Name *string `json:"name,omitempty"`

	// 委托权限信息
	Permissions *[]AgenciesPermissionInfo `json:"permissions,omitempty"`
}

func (o AgenciesInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgenciesInfo struct{}"
	}

	return strings.Join([]string{"AgenciesInfo", string(data)}, " ")
}
