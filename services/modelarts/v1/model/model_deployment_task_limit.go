package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeploymentTaskLimit struct {

	// **参数解释：** 最大并发任务数 **约束限制：** 不填保留原有值。 **取值范围：** [1, 100]。 **默认取值：** 不涉及
	MaxConcurrentTask *int32 `json:"max_concurrent_task,omitempty"`
}

func (o DeploymentTaskLimit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeploymentTaskLimit struct{}"
	}

	return strings.Join([]string{"DeploymentTaskLimit", string(data)}, " ")
}
