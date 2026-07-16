package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskAlgorithmRemote **参数解释**：数据实际输出信息。 **约束限制**：不涉及。
type TaskAlgorithmRemote struct {
	Obs *TaskAlgorithmRemoteObs `json:"obs"`
}

func (o TaskAlgorithmRemote) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskAlgorithmRemote struct{}"
	}

	return strings.Join([]string{"TaskAlgorithmRemote", string(data)}, " ")
}
