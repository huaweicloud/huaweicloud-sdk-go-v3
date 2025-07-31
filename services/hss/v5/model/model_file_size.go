package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FileSize **参数解释**: 文件大小 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值9223372036854775807 **默认取值**: 不涉及
type FileSize struct {
}

func (o FileSize) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileSize struct{}"
	}

	return strings.Join([]string{"FileSize", string(data)}, " ")
}
