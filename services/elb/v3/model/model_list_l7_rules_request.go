package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListL7RulesRequest Request Object
type ListL7RulesRequest struct {

	// **参数解释**：策略ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	L7policyId string `json:"l7policy_id"`

	// **参数解释**：每页返回的个数。  **约束限制**：不涉及  **取值范围**：0-2000  **默认取值**：2000
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：上一页最后一条记录的ID。  **约束限制**： - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。  **取值范围**：不涉及  **默认取值**：不涉及
	Marker *string `json:"marker,omitempty"`

	// **参数解释**：是否反向查询。  **约束限制**： - 必须与limit一起使用。 - 当page_reverse=true时，若要查询上一页，marker取值为当前页返回值的previous_marker。  **取值范围**： - true：查询上一页。 - false：查询下一页。  **默认取值**：false
	PageReverse *bool `json:"page_reverse,omitempty"`

	// **参数解释**：转发规则ID。 支持多值查询，查询条件格式：*id=xxx&id=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Id *[]string `json:"id,omitempty"`

	// **参数解释**：转发匹配方式。 支持多值查询，查询条件格式：*compare_type=xxx&compare_type=xxx*。  **约束限制**：不涉及  **取值范围**： - EQUAL_TO 表示精确匹配。 - REGEX 表示正则匹配。 - STARTS_WITH 表示前缀匹配。  **默认取值**：不涉及
	CompareType *[]string `json:"compare_type,omitempty"`

	// **参数解释**：转发规则的配置状态。 支持多值查询，查询条件格式：*provisioning_status=xxx&provisioning_status=xxx*。  **约束限制**：不涉及  **取值范围**：ACTIVE 表示正常。  **默认取值**：不涉及
	ProvisioningStatus *[]string `json:"provisioning_status,omitempty"`

	// **参数解释**：是否反向匹配。  **约束限制**：不涉及  **取值范围**：false  **默认取值**：不涉及
	Invert *bool `json:"invert,omitempty"`

	// **参数解释**：转发规则的管理状态。 不支持该字段，请勿使用。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// **参数解释**：匹配内容的值。 支持多值查询，查询条件格式：*value=xxx&value=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Value *[]string `json:"value,omitempty"`

	// **参数解释**：匹配内容的键值，用于标识规则。 支持多值查询，查询条件格式：*key=xxx&key=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Key *[]string `json:"key,omitempty"`

	// **参数解释**：匹配类别，可以为HOST_NAME，PATH。 支持多值查询，查询条件格式：*type=xxx&type=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Type *[]string `json:"type,omitempty"`

	// **参数解释**：资源所属的企业项目ID。 支持多值查询，查询条件格式：*enterprise_project_id=xxx&enterprise_project_id=xxx*。  **约束限制**： - 如果enterprise_project_id不传值，默认查询所有企业项目下的资源，鉴权按照细粒度权限鉴权，必须在用户组下分配elb:l7rules:list权限。 - 如果enterprise_project_id传值，鉴权按照企业项目权限鉴权，分为传入具体eps_id和all_granted_eps两种场景，前者查询指定eps_id的eps下的资源，后者查询的是所有有list权限的eps下的资源。  **取值范围**：不涉及  **默认取值**：不涉及  [不支持该字段，请勿使用。](tag:dt,hcso_dt)
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
}

func (o ListL7RulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListL7RulesRequest struct{}"
	}

	return strings.Join([]string{"ListL7RulesRequest", string(data)}, " ")
}
