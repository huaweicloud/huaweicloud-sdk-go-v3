package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskSettingsResponse Response Object
type ShowTaskSettingsResponse struct {

	// **参数解释：** 高级配置的相关信息。
	Info *[]TaskAdvancedSettings `json:"info,omitempty"`

	// **参数解释：** 总数。 **取值范围：** ≥0
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowTaskSettingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskSettingsResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskSettingsResponse", string(data)}, " ")
}
