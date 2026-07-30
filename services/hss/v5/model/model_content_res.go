package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ContentRes **参数解释**: 策略详情 **取值范围**: 字符长度0-65535位
type ContentRes struct {
}

func (o ContentRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContentRes struct{}"
	}

	return strings.Join([]string{"ContentRes", string(data)}, " ")
}
