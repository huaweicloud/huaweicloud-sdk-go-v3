package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RebuildRepositoryNavigationResponse Response Object
type RebuildRepositoryNavigationResponse struct {

	// **参数解释：** 结果标识。 **约束限制：** 不涉及。
	Result *string `json:"result,omitempty"`

	// **参数解释：** 结果消息。 **约束限制：** 不涉及。
	Message *string `json:"message,omitempty"`

	// **参数解释：** 触发任务耗时（毫秒）。 **约束限制：** 不涉及。
	Duration *string `json:"duration,omitempty"`

	// **参数解释：** 当前代码导航索引大小（字节）。 **约束限制：** 不涉及。
	Size           *int64 `json:"size,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o RebuildRepositoryNavigationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RebuildRepositoryNavigationResponse struct{}"
	}

	return strings.Join([]string{"RebuildRepositoryNavigationResponse", string(data)}, " ")
}
