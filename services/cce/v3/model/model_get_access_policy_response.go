package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetAccessPolicyResponse Response Object
type GetAccessPolicyResponse struct {

	// **参数解释：** API类型。 **约束限制：** 该值不可修改。 **取值范围：** 不涉及 **默认取值：** AccessPolicy
	Kind *string `json:"kind,omitempty"`

	// **参数解释：** API版本。 **约束限制：** 该值不可修改。 **取值范围：** 不涉及 **默认取值：** v3
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释：** 访问策略名称。 **约束限制：** 以小写字母开头，由小写字母、数字、中划线(-)、点(.)组成，长度范围1-56位，且不能以中划线(-)结尾。 **取值范围：** 不涉及 **默认取值：** 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释：** 权限ID。 **约束限制：** 系统自动生成，该值不可修改。 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyId *string `json:"policyId,omitempty"`

	// **参数解释：** 集群ID的列表，允许使用通配符（“\\*”），表示所有集群。 **约束限制：** 当前最多支持同时授权200个集群。 **取值范围：** \\[\"\\*\"\\]或者集群ID列表。 **默认取值：** 不涉及
	Clusters *[]string `json:"clusters,omitempty"`

	AccessScope *AccessScope `json:"accessScope,omitempty"`

	// **参数解释：** 权限类型。 **约束限制：** 不涉及 **取值范围：** - CCEClusterAdminPolicy：管理员权限 - CCEAdminPolicy：运维权限 - CCEEditPolicy：开发权限 - CCEViewPolicy：只读权限  **默认取值：** 不涉及
	PolicyType *string `json:"policyType,omitempty"`

	Principal *Principal `json:"principal,omitempty"`

	// **参数解释：** 创建时间。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	CreateTime *sdktime.SdkTime `json:"createTime,omitempty"`

	// **参数解释：** 更新时间。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	UpdateTime     *sdktime.SdkTime `json:"updateTime,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o GetAccessPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetAccessPolicyResponse struct{}"
	}

	return strings.Join([]string{"GetAccessPolicyResponse", string(data)}, " ")
}
