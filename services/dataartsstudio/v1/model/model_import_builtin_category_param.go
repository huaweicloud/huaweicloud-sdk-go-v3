package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImportBuiltinCategoryParam struct {

	// 规则对应密级的列表，需要将所有未导入的内置规则导入。
	RuleSecrecyLevelList *[]ImportRuleSecrecyLevelDto `json:"rule_secrecy_level_list,omitempty"`
}

func (o ImportBuiltinCategoryParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportBuiltinCategoryParam struct{}"
	}

	return strings.Join([]string{"ImportBuiltinCategoryParam", string(data)}, " ")
}
