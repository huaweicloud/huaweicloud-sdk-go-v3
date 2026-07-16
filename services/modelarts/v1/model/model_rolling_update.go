package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RollingUpdate **参数解释：** 滚动更新配置。 **约束限制：** 不涉及。
type RollingUpdate struct {

	// **参数解释：** 滚动更新时最多可以启动多少个Pod。 **约束限制：** 百分数类型字符串。 **取值范围：** 1%-100%。 **默认取值：** 1%。
	MaxSurge *string `json:"max_surge,omitempty"`

	// **参数解释：** 滚动更新时最多可以删除多少个Pod。 **约束限制：** 百分数类型字符串。 **取值范围：** 1%-100%。 **默认取值：** 1%。
	MaxUnavailable *string `json:"max_unavailable,omitempty"`
}

func (o RollingUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollingUpdate struct{}"
	}

	return strings.Join([]string{"RollingUpdate", string(data)}, " ")
}
