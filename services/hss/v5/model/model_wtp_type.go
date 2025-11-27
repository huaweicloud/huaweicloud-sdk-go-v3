package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WtpType **参数解释**: 网页防篡改的类型 **约束限制**: 不涉及 **取值范围**: - container_wtp：容器网页防篡改  **默认取值**: 不涉及
type WtpType struct {
}

func (o WtpType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WtpType struct{}"
	}

	return strings.Join([]string{"WtpType", string(data)}, " ")
}
