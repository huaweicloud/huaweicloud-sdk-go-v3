package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataInfo **参数解释**：模型信息（OBS 输入路径）。RayService 相对 RayCluster 新增的核心配置，仅需指明模型在 OBS 的位置。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
type DataInfo struct {

	// **参数解释**：OBS 模型输入路径（模型在 OBS 的位置）。 **约束限制**：不涉及。 **取值范围**：长度为1~1024的字符。 **默认取值**：不涉及。
	InputPath string `json:"input_path"`
}

func (o DataInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataInfo struct{}"
	}

	return strings.Join([]string{"DataInfo", string(data)}, " ")
}
