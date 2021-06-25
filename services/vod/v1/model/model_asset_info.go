package model

import (
	"encoding/json"

	"strings"
)

type AssetInfo struct {
	// 媒体ID<br/>

	AssetId *string `json:"asset_id,omitempty"`

	Status *Status `json:"status,omitempty"`
	// 媒资描述信息<br/>

	Description *string `json:"description,omitempty"`

	BaseInfo *BaseInfo `json:"base_info,omitempty"`
	// 播放信息数组<br/>

	PlayInfoArray *[]PlayInfo `json:"play_info_array,omitempty"`
}

func (o AssetInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AssetInfo struct{}"
	}

	return strings.Join([]string{"AssetInfo", string(data)}, " ")
}
