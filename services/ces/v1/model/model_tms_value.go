package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TmsValue **参数解释** TMS标签值规范 **约束限制** 不涉及 **取值范围** 长度为[0，43]个字符 **默认取值** 不涉及
type TmsValue struct {
}

func (o TmsValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TmsValue struct{}"
	}

	return strings.Join([]string{"TmsValue", string(data)}, " ")
}
