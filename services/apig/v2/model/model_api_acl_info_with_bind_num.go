package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiAclInfoWithBindNum struct {
	// ACL策略名称

	AclName *string `json:"acl_name,omitempty"`
	// 类型 - PERMIT（白名单类型） - DENY（黑名单类型）

	AclType *string `json:"acl_type,omitempty"`
	// ACL策略的值

	AclValue *string `json:"acl_value,omitempty"`
	// 绑定的API数量

	BindNum *int32 `json:"bind_num,omitempty"`
	// 对象类型 - IP - DOMAIN

	EntityType *string `json:"entity_type,omitempty"`
	// ACL策略编号

	Id *string `json:"id,omitempty"`
	// 更新时间

	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
}

func (o ApiAclInfoWithBindNum) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiAclInfoWithBindNum struct{}"
	}

	return strings.Join([]string{"ApiAclInfoWithBindNum", string(data)}, " ")
}
