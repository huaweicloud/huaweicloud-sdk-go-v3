package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisplayNameReq **参数解释**: 主题的显示名 **约束限制**: 不涉及 **取值范围**: 字符长度0-256位 **默认取值**: 不涉及
type DisplayNameReq struct {
}

func (o DisplayNameReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisplayNameReq struct{}"
	}

	return strings.Join([]string{"DisplayNameReq", string(data)}, " ")
}
