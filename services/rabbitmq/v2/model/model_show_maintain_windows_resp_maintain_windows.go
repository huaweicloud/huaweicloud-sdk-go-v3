package model

import (
	"encoding/json"

	"strings"
)

type ShowMaintainWindowsRespMaintainWindows struct {
	// 是否为默认时间段。

	Default *bool `json:"default,omitempty"`
	// 维护时间窗结束时间。

	End *string `json:"end,omitempty"`
	// 维护时间窗开始时间。

	Begin *string `json:"begin,omitempty"`
	// 序号。

	Seq *int32 `json:"seq,omitempty"`
}

func (o ShowMaintainWindowsRespMaintainWindows) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowMaintainWindowsRespMaintainWindows struct{}"
	}

	return strings.Join([]string{"ShowMaintainWindowsRespMaintainWindows", string(data)}, " ")
}
