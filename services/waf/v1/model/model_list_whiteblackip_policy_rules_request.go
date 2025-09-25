package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWhiteblackipPolicyRulesRequest Request Object
type ListWhiteblackipPolicyRulesRequest struct {

	// **参数解释：** 项目ID，对应控制台用户名->我的凭证->项目列表->项目ID。 **约束限制：** 不涉及 **取值范围：** 只能由英文字母、数字组成，且长度为32个字符。 **默认取值：** 不涉及
	Projectid string `json:"projectid"`

	// **参数解释：** 域名id。 **约束限制：** 不涉及 **取值范围：** 只能由英文字母、数字组成，且长度为32个字符。 **默认取值：** 不涉及
	Policyids *string `json:"policyids,omitempty"`

	// **参数解释：** 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目ID。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。 **约束限制：** 不涉及 **取值范围：**  - 0：代表default企业项目  - all_granted_eps：代表所有企业项目  - 其它企业项目ID：长度为36个字符  **默认取值：** 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释：** 分页查询的起始位置，表示从第几条记录开始返回（从1开始计数）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 1
	Page *int32 `json:"page,omitempty"`

	// **参数解释：** 分页查询时，每页包含多少条结果。 **约束限制：** 不涉及 **默认取值：** 1000
	Pagesize *int32 `json:"pagesize,omitempty"`
}

func (o ListWhiteblackipPolicyRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWhiteblackipPolicyRulesRequest struct{}"
	}

	return strings.Join([]string{"ListWhiteblackipPolicyRulesRequest", string(data)}, " ")
}
