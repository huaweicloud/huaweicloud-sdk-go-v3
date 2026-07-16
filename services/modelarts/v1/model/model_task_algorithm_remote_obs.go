package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskAlgorithmRemoteObs **参数解释**：数据实际输出到OBS。 **约束限制**：不涉及。
type TaskAlgorithmRemoteObs struct {

	// **参数解释**：数据实际输出到OBS的路径。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	ObsUrl string `json:"obs_url"`
}

func (o TaskAlgorithmRemoteObs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskAlgorithmRemoteObs struct{}"
	}

	return strings.Join([]string{"TaskAlgorithmRemoteObs", string(data)}, " ")
}
