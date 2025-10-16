package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateIgnoreRuleRequest Request Object
type BatchCreateIgnoreRuleRequest struct {

	// **参数解释：** 项目ID，对应控制台用户名->我的凭证->项目列表->项目ID。 **约束限制：** 不涉及 **取值范围：** 只能由英文字母、数字组成，且长度为32个字符。 **默认取值：** 不涉及
	Projectid string `json:"projectid"`

	// **参数解释：** 域名id。 **约束限制：** 不涉及 **取值范围：** 只能由英文字母、数字组成，且长度为32个字符。 **默认取值：** 不涉及
	Policyids *string `json:"policyids,omitempty"`

	// **参数解释：** 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目ID。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。 **约束限制：** 不涉及 **取值范围：**  - 0：代表default企业项目  - all_granted_eps：代表所有企业项目  - 其它企业项目ID：长度为36个字符  **默认取值：** 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *BatchCreateIgnoreRuleRequestBody `json:"body,omitempty"`
}

func (o BatchCreateIgnoreRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateIgnoreRuleRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateIgnoreRuleRequest", string(data)}, " ")
}
