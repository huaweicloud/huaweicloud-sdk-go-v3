package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogoTypeRes **参数解释**: 报告logo类型 **取值范围**: - default：华为logo - customize：自定义logo
type LogoTypeRes struct {
}

func (o LogoTypeRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogoTypeRes struct{}"
	}

	return strings.Join([]string{"LogoTypeRes", string(data)}, " ")
}
