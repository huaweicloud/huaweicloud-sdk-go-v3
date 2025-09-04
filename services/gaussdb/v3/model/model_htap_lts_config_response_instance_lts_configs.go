package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HtapLtsConfigResponseInstanceLtsConfigs struct {

	// **参数解释**： LTS配置信息。  **约束限制**： 不涉及。
	LtsConfigs []HtapLtsConfigResponseLtsConfigs `json:"lts_configs"`

	Instance *HtapLtsConfigResponseInstance `json:"instance"`
}

func (o HtapLtsConfigResponseInstanceLtsConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HtapLtsConfigResponseInstanceLtsConfigs struct{}"
	}

	return strings.Join([]string{"HtapLtsConfigResponseInstanceLtsConfigs", string(data)}, " ")
}
