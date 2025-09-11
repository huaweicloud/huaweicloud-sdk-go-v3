package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TmsKey **参数解释** TMS标签键规范 **约束限制** 不涉及 **取值范围** 长度为[1,128]个字符 **默认取值** 不涉及
type TmsKey struct {
}

func (o TmsKey) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TmsKey struct{}"
	}

	return strings.Join([]string{"TmsKey", string(data)}, " ")
}
