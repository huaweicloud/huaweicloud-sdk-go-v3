package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UserPolicyResponeEntity struct {

	// **参数解释**： 资源类型。 **取值范围**： - TOPIC：表示资源类型为Topic。
	ResourceType *string `json:"resource_type,omitempty"`

	// **参数解释**： 资源名称。 **取值范围**： - 已有的Topic名称。 - 符合Topic名称规则的前缀。 - “*”表示匹配所有的Topic。
	ResourceName *string `json:"resource_name,omitempty"`

	// **参数解释**： 匹配方式。 **取值范围**： - LITERAL：完全匹配。 - PREFIXED：前缀匹配。
	PatternType *string `json:"pattern_type,omitempty"`

	// **参数解释**： 权限类型。 **取值范围**： - all：拥有发布、订阅权限。 - pub：拥有发布权限。 - sub：拥有订阅权限。
	AccessPolicy *string `json:"access_policy,omitempty"`

	// **参数解释**： Acl权限类型。 **取值范围**： - ALLOW：允许用户进行某种操作。
	AclPermissionType *string `json:"acl_permission_type,omitempty"`
}

func (o UserPolicyResponeEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserPolicyResponeEntity struct{}"
	}

	return strings.Join([]string{"UserPolicyResponeEntity", string(data)}, " ")
}
