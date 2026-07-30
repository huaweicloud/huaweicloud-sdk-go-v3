package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ContentReqM **参数解释**： 策略详情 **约束限制**： 必填 **取值范围**： 字符长度1-65535位 **默认取值**： 不涉及
type ContentReqM struct {
}

func (o ContentReqM) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContentReqM struct{}"
	}

	return strings.Join([]string{"ContentReqM", string(data)}, " ")
}
