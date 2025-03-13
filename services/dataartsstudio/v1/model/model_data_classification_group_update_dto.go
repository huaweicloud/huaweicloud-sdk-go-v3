package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataClassificationGroupUpdateDto struct {

	// 规则名称
	Name string `json:"name"`

	// 规则id列表
	RuleIds []string `json:"rule_ids"`

	// 规则组描述
	Description *string `json:"description,omitempty"`
}

func (o DataClassificationGroupUpdateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataClassificationGroupUpdateDto struct{}"
	}

	return strings.Join([]string{"DataClassificationGroupUpdateDto", string(data)}, " ")
}
