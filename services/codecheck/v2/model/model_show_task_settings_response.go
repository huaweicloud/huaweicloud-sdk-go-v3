package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowTaskSettingsResponse struct {

	// 高级选项的相关信息
	Info *[]TaskAdvancedSettings `json:"info,omitempty"`

	// 总数
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
