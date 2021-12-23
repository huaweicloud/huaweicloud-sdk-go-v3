package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AclApiBindingInfo struct {
	// 绑定关系编号

	Id *string `json:"id,omitempty"`
	// API编号

	ApiId *string `json:"api_id,omitempty"`
	// 环境编号

	EnvId *string `json:"env_id,omitempty"`
	// ACL策略编号

	AclId *string `json:"acl_id,omitempty"`
	// 绑定时间

	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`
}

func (o AclApiBindingInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AclApiBindingInfo struct{}"
	}

	return strings.Join([]string{"AclApiBindingInfo", string(data)}, " ")
}
