package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateWhiteblackipRuleRequest Request Object
type BatchCreateWhiteblackipRuleRequest struct {

	// **参数解释：** 策略id列表。策略id从\"查询防护策略列表\"(ListPolicy)接口获取，多个策略之间用“,”隔开 **约束限制：** 不涉及 **取值范围：** 策略id只能由英文字母、数字组成，且长度为32个字符。 **默认取值：** 不涉及
	Policyids *string `json:"policyids,omitempty"`

	// **参数解释：** 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目ID。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。 **约束限制：** 不涉及 **取值范围：**  - 0：代表default企业项目  - all_granted_eps：代表所有企业项目  - 其它企业项目ID：长度为36个字符  **默认取值：** 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *BacthCreateWhiteBlackIpRuleRequestBody `json:"body,omitempty"`
}

func (o BatchCreateWhiteblackipRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateWhiteblackipRuleRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateWhiteblackipRuleRequest", string(data)}, " ")
}
