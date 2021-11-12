package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListL7RulesRequest struct {
	// 策略ID。

	L7policyId string `json:"l7policy_id"`
	// 转发规则的管理状态；该字段为预留字段，暂未启用。默认为true。

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 转发规则的匹配方式。type为HOST_NAME时可以为EQUAL_TO。type为PATH时可以为REGEX， STARTS_WITH，EQUAL_TO。

	CompareType *[]string `json:"compare_type,omitempty"`
	// 企业项目ID。

	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
	// 转发规则ID。

	Id *[]string `json:"id,omitempty"`
	// 是否反向匹配。使用说明：固定为false。该字段能更新但不会生效。

	Invert *bool `json:"invert,omitempty"`
	// 匹配内容的键值。目前匹配内容为HOST_NAME和PATH时，该字段不生效。该字段能更新但不会生效。

	Key *[]string `json:"key,omitempty"`
	// 每页返回的个数。

	Limit *int32 `json:"limit,omitempty"`
	// 上一页最后一条记录的ID。  使用说明：  - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。

	Marker *string `json:"marker,omitempty"`
	// 分页的顺序，true表示从后往前分页，false表示从前往后分页，默认为false。 使用说明：必须与limit一起使用。

	PageReverse *bool `json:"page_reverse,omitempty"`
	// 转发规则的配置状态；该字段为预留字段，暂未启用。默认为ACTIVE。

	ProvisioningStatus *[]string `json:"provisioning_status,omitempty"`
	// 一个l7policy下创建的l7rule的type不能重复。 匹配内容：可以为HOST_NAME，PATH。

	Type *[]string `json:"type,omitempty"`
	// 匹配内容的值。其值不能包含空格。使用说明：当type为HOST_NAME时，取值范围：String(100)，字符串只能包含英文字母、数字、“-”或“.”，且必须以字母或数字开头。当type为PATH时，取值范围：String(128)。当转发规则的compare_type为STARTS_WITH，EQUAL_TO时，字符串只能包含英文字母、数字、^-%#&$.*+?,=!:| /()[]{}，且必须以\"/\"开头。

	Value *[]string `json:"value,omitempty"`
}

func (o ListL7RulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListL7RulesRequest struct{}"
	}

	return strings.Join([]string{"ListL7RulesRequest", string(data)}, " ")
}
