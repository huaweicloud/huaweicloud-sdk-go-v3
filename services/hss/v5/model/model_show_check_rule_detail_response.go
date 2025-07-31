package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCheckRuleDetailResponse Response Object
type ShowCheckRuleDetailResponse struct {

	// **参数解释**: 当前检查项（检测规则）的描述 **取值范围**: 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**: 当前检查项（检测规则）的制定依据 **取值范围**: 不涉及
	Reference *string `json:"reference,omitempty"`

	// **参数解释**: 当前检查项（检测规则）的审计描述 **取值范围**: 不涉及
	Audit *string `json:"audit,omitempty"`

	// **参数解释**: 当前检查项（检测规则）的修改建议 **取值范围**: 不涉及
	Remediation *string `json:"remediation,omitempty"`

	// **参数解释**: 检测用例信息 **取值范围**: 不涉及
	CheckInfoList  *[]CheckRuleCheckCaseResponseInfo `json:"check_info_list,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ShowCheckRuleDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCheckRuleDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowCheckRuleDetailResponse", string(data)}, " ")
}
