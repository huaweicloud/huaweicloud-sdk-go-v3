package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CongestionLanesInfo 车道拥堵信息。
type CongestionLanesInfo struct {

	// **参数说明**：车道号。
	Laneid *int32 `json:"laneid,omitempty"`

	// **参数说明**：拥堵级别。  **取值范围**：  - 1：拥堵级别低，速度[25, 30) 单位：km/h  - 2：拥堵级别中，速度[15，25) 单位：km/h  - 3：拥堵级别高，速度[0, 15) 单位：km/h
	Level *int32 `json:"level,omitempty"`

	// **参数说明**：拥堵长度，单位为米（m）。
	Length *int32 `json:"length,omitempty"`

	StartPoint *Position3D `json:"start_point,omitempty"`

	EndPoint *Position3D `json:"end_point,omitempty"`
}

func (o CongestionLanesInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CongestionLanesInfo struct{}"
	}

	return strings.Join([]string{"CongestionLanesInfo", string(data)}, " ")
}
