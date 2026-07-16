package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RollingUpdateResponse **参数解释：** 滚动更新配置。
type RollingUpdateResponse struct {

	// **参数解释：** 滚动更新时最多可以启动多少个Pod。 **取值范围：** 不涉及。
	MaxSurge *string `json:"max_surge,omitempty"`

	// **参数解释：** 滚动更新时最多可以删除多少个Pod。 **取值范围：** 不涉及。
	MaxUnavailable *string `json:"max_unavailable,omitempty"`
}

func (o RollingUpdateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollingUpdateResponse struct{}"
	}

	return strings.Join([]string{"RollingUpdateResponse", string(data)}, " ")
}
