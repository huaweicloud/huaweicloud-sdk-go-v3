package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataClassificationRuleEnableDto struct {

	// 是否开启识别规则
	Enable *bool `json:"enable,omitempty"`
}

func (o DataClassificationRuleEnableDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataClassificationRuleEnableDto struct{}"
	}

	return strings.Join([]string{"DataClassificationRuleEnableDto", string(data)}, " ")
}
