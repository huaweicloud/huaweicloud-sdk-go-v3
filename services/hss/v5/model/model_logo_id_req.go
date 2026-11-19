package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogoIdReq **参数解释**: 自定义logo的文件id **约束限制**: 不涉及 **取值范围**: 字符长度0-256位 **默认取值**: 不涉及
type LogoIdReq struct {
}

func (o LogoIdReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogoIdReq struct{}"
	}

	return strings.Join([]string{"LogoIdReq", string(data)}, " ")
}
