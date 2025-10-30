package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PwdCheckSubTagInfo 口令策略二级标签信息
type PwdCheckSubTagInfo struct {

	// **参数解释** 口令检测二级tag **取值范围** - weakpwd - pwdcomplexity
	SubTag *string `json:"sub_tag,omitempty"`

	// **参数解释** 检测范围。 **取值范围** - true: 如果sub_tag有值,则二级标签sub_tag下所有的检测项都会检测 - false: 如果sub_tag有值，则二级标签sub_tag下所有的检测项都不会检测。 - indeterminate: 部分检测，由check_rule_ids决定具体检测项。
	Checked *string `json:"checked,omitempty"`

	// 检测项列表
	CheckRuleIds *[]string `json:"check_rule_ids,omitempty"`
}

func (o PwdCheckSubTagInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PwdCheckSubTagInfo struct{}"
	}

	return strings.Join([]string{"PwdCheckSubTagInfo", string(data)}, " ")
}
