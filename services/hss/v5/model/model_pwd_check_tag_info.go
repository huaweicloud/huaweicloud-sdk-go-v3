package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PwdCheckTagInfo 口令策略一级标签信息
type PwdCheckTagInfo struct {

	// **参数解释** 口令检测一级tag **取值范围** - weakpwd_pwdcomplexity
	Tag *string `json:"tag,omitempty"`

	// **参数解释** 检测范围。 **取值范围** - true: 如果tag有值,则一级标签tag下所有的检测项都会检测。 - false: 如果tag有值，则一级标签tag下所有的检测项都不会检测。 - indeterminate: 部分检测，具体检测项由sub_tags决定。
	Checked *string `json:"checked,omitempty"`

	// 服务器风险TOP5列表
	SubTags *[]PwdCheckSubTagInfo `json:"sub_tags,omitempty"`
}

func (o PwdCheckTagInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PwdCheckTagInfo struct{}"
	}

	return strings.Join([]string{"PwdCheckTagInfo", string(data)}, " ")
}
