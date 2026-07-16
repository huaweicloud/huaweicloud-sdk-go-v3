package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Obs 训练作业obs挂载信息。
type Obs struct {

	// **参数解释**：需要挂载的obs路径。如：“/test-bucket/path”。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	ObsPath *string `json:"obs_path,omitempty"`

	// **参数解释**：挂载到训练容器中的路径，如：“/example/path”。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	LocalPath *string `json:"local_path,omitempty"`
}

func (o Obs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Obs struct{}"
	}

	return strings.Join([]string{"Obs", string(data)}, " ")
}
