package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHtapLtsConfigResponse Response Object
type ShowHtapLtsConfigResponse struct {

	// **参数解释**： 实例LTS配置信息。  **约束限制**： 不涉及。
	InstanceLtsConfigs *[]HtapLtsConfigResponseInstanceLtsConfigs `json:"instance_lts_configs,omitempty"`

	// **参数解释**： 实例数。  **取值范围**：  不涉及。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowHtapLtsConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHtapLtsConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowHtapLtsConfigResponse", string(data)}, " ")
}
