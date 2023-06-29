package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CongestionInfo 拥堵事件信息
type CongestionInfo struct {

	// **参数说明**：拥堵级别。  **取值范围**：  - 1：拥堵级别低，速度[25, 30) 单位：km/h  - 2：拥堵级别中，速度[15，25) 单位：km/h  - 3：拥堵级别高，速度[0, 15) 单位：km/h
	Level *int32 `json:"level,omitempty"`

	// **参数说明**：拥堵长度，单位为米（m）。
	Length *int32 `json:"length,omitempty"`

	StartPoint *Position3D `json:"start_point,omitempty"`

	EndPoint *Position3D `json:"end_point,omitempty"`

	// **参数说明**：车道拥堵信息。
	CongestionLanesInfo *[]CongestionLanesInfo `json:"congestion_lanes_info,omitempty"`
}

func (o CongestionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CongestionInfo struct{}"
	}

	return strings.Join([]string{"CongestionInfo", string(data)}, " ")
}
