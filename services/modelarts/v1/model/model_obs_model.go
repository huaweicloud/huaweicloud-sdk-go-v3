package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ObsModel **参数解释**：自定义训练作业产物obs输出信息。
type ObsModel struct {

	// **参数解释**：自定义训练作业产物保存的OBS地址，如：“obs://example/path”。 **取值范围**：不涉及。
	ObsPath string `json:"obs_path"`

	// **参数解释**：自定义训练作业产物保存的宿主机的路径，如：“/example/path”。 **取值范围**：不涉及。
	LocalPath *string `json:"local_path,omitempty"`
}

func (o ObsModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsModel struct{}"
	}

	return strings.Join([]string{"ObsModel", string(data)}, " ")
}
