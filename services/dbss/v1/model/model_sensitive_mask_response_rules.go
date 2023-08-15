package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SensitiveMaskResponseRules struct {

	// 规则ID
	Id *string `json:"id,omitempty"`

	// 规则名称
	Name *string `json:"name,omitempty"`

	// 规则类型
	Type *string `json:"type,omitempty"`

	// 规则正则表达式
	Regex *string `json:"regex,omitempty"`

	// 替换值
	MaskValue *string `json:"mask_value,omitempty"`

	// 规则状态
	Status *string `json:"status,omitempty"`

	// 操作时间
	OperateTime *string `json:"operate_time,omitempty"`
}

func (o SensitiveMaskResponseRules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SensitiveMaskResponseRules struct{}"
	}

	return strings.Join([]string{"SensitiveMaskResponseRules", string(data)}, " ")
}
