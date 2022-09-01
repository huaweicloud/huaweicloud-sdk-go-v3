package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateAclStrategyV2Response struct {

	// 名称
	AclName *string `json:"acl_name,omitempty" xml:"acl_name"`

	// 类型: - PERMIT（白名单类型） - DENY（黑名单类型）
	AclType *string `json:"acl_type,omitempty" xml:"acl_type"`

	// ACL策略值
	AclValue *string `json:"acl_value,omitempty" xml:"acl_value"`

	// 对象类型： - IP - DOMAIN
	EntityType *string `json:"entity_type,omitempty" xml:"entity_type"`

	// 编号
	Id *string `json:"id,omitempty" xml:"id"`

	// 更新时间
	UpdateTime     *sdktime.SdkTime `json:"update_time,omitempty" xml:"update_time"`
	HttpStatusCode int              `json:"-"`
}

func (o UpdateAclStrategyV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAclStrategyV2Response struct{}"
	}

	return strings.Join([]string{"UpdateAclStrategyV2Response", string(data)}, " ")
}
