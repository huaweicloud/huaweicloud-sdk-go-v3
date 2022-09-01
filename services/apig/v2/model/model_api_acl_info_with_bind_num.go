package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiAclInfoWithBindNum struct {

	// ACL策略名称
	AclName *string `json:"acl_name,omitempty" xml:"acl_name"`

	// 类型 - PERMIT（白名单类型） - DENY（黑名单类型）
	AclType *string `json:"acl_type,omitempty" xml:"acl_type"`

	// ACL策略的值
	AclValue *string `json:"acl_value,omitempty" xml:"acl_value"`

	// 绑定的API数量
	BindNum *int32 `json:"bind_num,omitempty" xml:"bind_num"`

	// 对象类型 - IP - DOMAIN
	EntityType *string `json:"entity_type,omitempty" xml:"entity_type"`

	// ACL策略编号
	Id *string `json:"id,omitempty" xml:"id"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty" xml:"update_time"`
}

func (o ApiAclInfoWithBindNum) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiAclInfoWithBindNum struct{}"
	}

	return strings.Join([]string{"ApiAclInfoWithBindNum", string(data)}, " ")
}
