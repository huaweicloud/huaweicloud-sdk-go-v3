package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListL7RulesRequest struct {
	// 策略ID。

	L7policyId string `json:"l7policy_id"`
	// 每页返回的个数。

	Limit *int32 `json:"limit,omitempty"`
	// 上一页最后一条记录的ID。  使用说明：  - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。

	Marker *string `json:"marker,omitempty"`
	// 分页的顺序，true表示从后往前分页，false表示从前往后分页，默认为false。  使用说明： - 必须与limit一起使用。

	PageReverse *bool `json:"page_reverse,omitempty"`
	// 转发规则ID。  支持多值查询，查询条件格式：*id=xxx&id=xxx*。

	Id *[]string `json:"id,omitempty"`
	// 转发规则的匹配方式。  type为HOST_NAME时可以为EQUAL_TO。  type为PATH时可以为Perl类型的REGEX， STARTS_WITH，EQUAL_TO。  支持多值查询，查询条件格式：*compare_type=xxx&compare_type=xxx*。

	CompareType *[]string `json:"compare_type,omitempty"`
	// 转发规则的配置状态。取值：ACTIVE 表示正常。  支持多值查询，查询条件格式：*provisioning_status=xxx&provisioning_status=xxx*。

	ProvisioningStatus *[]string `json:"provisioning_status,omitempty"`
	// 是否反向匹配。固定为false。该字段能更新但不会生效。

	Invert *bool `json:"invert,omitempty"`
	// 转发规则的管理状态，默认为true。  不支持该字段，请勿使用。

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 匹配内容的值。  支持多值查询，查询条件格式：*value=xxx&value=xxx*。

	Value *[]string `json:"value,omitempty"`
	// 匹配内容的键值，用于标识规则。  支持多值查询，查询条件格式：*key=xxx&key=xxx*。  不支持该字段，请勿使用。

	Key *[]string `json:"key,omitempty"`
	// 匹配类别，可以为HOST_NAME，PATH。  一个l7policy下创建的l7rule的type不能重复。  支持多值查询，查询条件格式：*type=xxx&type=xxx*。

	Type *[]string `json:"type,omitempty"`
	// 企业项目ID。  支持多值查询，查询条件格式：*enterprise_project_id=xxx&enterprise_project_id=xxx*。  [不支持该字段，请勿使用。](tag:dt,dt_test,hcso_dt)

	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
}

func (o ListL7RulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListL7RulesRequest struct{}"
	}

	return strings.Join([]string{"ListL7RulesRequest", string(data)}, " ")
}
