package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogoIdRes **参数解释**: 自定义logo的文件id **取值范围**: 字符长度0-256位
type LogoIdRes struct {
}

func (o LogoIdRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogoIdRes struct{}"
	}

	return strings.Join([]string{"LogoIdRes", string(data)}, " ")
}
