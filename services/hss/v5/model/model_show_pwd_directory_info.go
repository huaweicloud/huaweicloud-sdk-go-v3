package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPwdDirectoryInfo 弱口令检查策略目录详细信息
type ShowPwdDirectoryInfo struct {

	// 弱口令及口令复杂度一级标签，包含如下:   - weakpwd_pwdcomplexity : 弱口令及口令复杂度检测   - weakpwd               : 弱口令检测
	Tag *string `json:"tag,omitempty"`

	// 口令检查包含子标签，包含如下:   - weak_pwd : 经典弱口令检测   - pwd_complexity : 口令复杂度策略检
	SubTag *string `json:"sub_tag,omitempty"`

	// **参数解释** 该项是否被选中 **取值范围** - true : 被选中 - false: 未被选中
	Checked *bool `json:"checked,omitempty"`

	// 表示目录中的唯一值:   - weak_pwd : 经典弱口令检测   - pwd_complexity : 口令复杂度策略检测
	Key *string `json:"key,omitempty"`
}

func (o ShowPwdDirectoryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPwdDirectoryInfo struct{}"
	}

	return strings.Join([]string{"ShowPwdDirectoryInfo", string(data)}, " ")
}
