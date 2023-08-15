package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDetailsOfAclPolicyV2Response Response Object
type ShowDetailsOfAclPolicyV2Response struct {

	// 名称
	AclName *string `json:"acl_name,omitempty"`

	// 类型: - PERMIT（白名单类型） - DENY（黑名单类型）
	AclType *string `json:"acl_type,omitempty"`

	// ACL策略值
	AclValue *string `json:"acl_value,omitempty"`

	// 对象类型： - IP - DOMAIN
	EntityType *string `json:"entity_type,omitempty"`

	// 编号
	Id *string `json:"id,omitempty"`

	// 更新时间
	UpdateTime     *sdktime.SdkTime `json:"update_time,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowDetailsOfAclPolicyV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailsOfAclPolicyV2Response struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfAclPolicyV2Response", string(data)}, " ")
}
