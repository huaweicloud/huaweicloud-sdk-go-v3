package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogoTypeReq **参数解释**: 报告logo类型 **约束限制**: 不涉及 **取值范围**: - default：华为logo - customize：自定义logo  **默认取值**: 不涉及
type LogoTypeReq struct {
}

func (o LogoTypeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogoTypeReq struct{}"
	}

	return strings.Join([]string{"LogoTypeReq", string(data)}, " ")
}
