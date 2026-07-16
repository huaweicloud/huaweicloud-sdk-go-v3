package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerTaskLimit **参数解释：** 服务任务限制信息。 **约束限制：** 不涉及。
type ServerTaskLimit struct {

	// **参数解释：** 单个服务任务限制总数。 **取值范围：** [0, 10000]。
	MaxTask *int32 `json:"max_task,omitempty"`
}

func (o ServerTaskLimit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerTaskLimit struct{}"
	}

	return strings.Join([]string{"ServerTaskLimit", string(data)}, " ")
}
