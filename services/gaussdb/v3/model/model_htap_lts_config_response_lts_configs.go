package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HtapLtsConfigResponseLtsConfigs struct {

	// **参数解释**： 日志类型。  **取值范围**：  不涉及。
	LogType string `json:"log_type"`

	// **参数解释**： LTS日志组ID。  **取值范围**：  不涉及。
	LtsGroupId string `json:"lts_group_id"`

	// **参数解释**： LTS日志流ID。  **取值范围**：  不涉及。
	LtsStreamId string `json:"lts_stream_id"`

	// **参数解释**： LTS配置开关状态。  **取值范围**：  不涉及。
	Enabled string `json:"enabled"`
}

func (o HtapLtsConfigResponseLtsConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HtapLtsConfigResponseLtsConfigs struct{}"
	}

	return strings.Join([]string{"HtapLtsConfigResponseLtsConfigs", string(data)}, " ")
}
