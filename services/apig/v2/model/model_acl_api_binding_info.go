package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AclApiBindingInfo struct {

	// 绑定关系编号
	Id *string `json:"id,omitempty" xml:"id"`

	// API编号
	ApiId *string `json:"api_id,omitempty" xml:"api_id"`

	// 环境编号
	EnvId *string `json:"env_id,omitempty" xml:"env_id"`

	// ACL策略编号
	AclId *string `json:"acl_id,omitempty" xml:"acl_id"`

	// 绑定时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty" xml:"create_time"`
}

func (o AclApiBindingInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AclApiBindingInfo struct{}"
	}

	return strings.Join([]string{"AclApiBindingInfo", string(data)}, " ")
}
