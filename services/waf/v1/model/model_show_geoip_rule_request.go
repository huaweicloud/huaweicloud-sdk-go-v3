package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGeoipRuleRequest Request Object
type ShowGeoipRuleRequest struct {

	// **参数解释：** 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目ID。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。 **约束限制：** 不涉及 **取值范围：**  - 0：代表default企业项目  - all_granted_eps：代表所有企业项目  - 其它企业项目ID：长度为36个字符  **默认取值：** 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防护策略id，您可以通过调用查询防护策略列表（ListPolicy）获取策略id
	PolicyId string `json:"policy_id"`

	// 地理位置控制防护规则id，规则id从查询地理位置控制防护规则列表（ListGeoipRule）接口获取，响应体的id字段
	RuleId string `json:"rule_id"`
}

func (o ShowGeoipRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGeoipRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowGeoipRuleRequest", string(data)}, " ")
}
