package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EndpointPermission struct {

	// 记录编号
	Id string `json:"id"`

	// 权限规则
	Permission string `json:"permission"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at"`
}

func (o EndpointPermission) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointPermission struct{}"
	}

	return strings.Join([]string{"EndpointPermission", string(data)}, " ")
}
