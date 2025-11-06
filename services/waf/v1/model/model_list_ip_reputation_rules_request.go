package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIpReputationRulesRequest Request Object
type ListIpReputationRulesRequest struct {

	// **参数解释：** 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目ID。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。 **约束限制：** 不涉及 **取值范围：**  - 0：代表default企业项目  - all_granted_eps：代表所有企业项目  - 其它企业项目ID：长度为36个字符 **默认取值：** 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释：** 策略id（策略id从查询防护策略列表接口获取） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyId string `json:"policy_id"`

	// **参数解释：** 偏移量，表示查询该偏移量之后的记录。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Offset int32 `json:"offset"`

	// **参数解释：** 查询返回记录的数量限制。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 1000
	Limit int32 `json:"limit"`
}

func (o ListIpReputationRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIpReputationRulesRequest struct{}"
	}

	return strings.Join([]string{"ListIpReputationRulesRequest", string(data)}, " ")
}
