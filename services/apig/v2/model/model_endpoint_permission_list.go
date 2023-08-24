package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EndpointPermissionList struct {

	// 权限列表
	Permissions []string `json:"permissions"`
}

func (o EndpointPermissionList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointPermissionList struct{}"
	}

	return strings.Join([]string{"EndpointPermissionList", string(data)}, " ")
}
