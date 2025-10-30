package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyWatermarkRule struct {

	// 水印规则名称，如果不填说明不修改。
	RuleName *string `json:"rule_name,omitempty"`

	Location *WatermarkLocation `json:"location"`
}

func (o ModifyWatermarkRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyWatermarkRule struct{}"
	}

	return strings.Join([]string{"ModifyWatermarkRule", string(data)}, " ")
}
