package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 分车道统计的占有率
type LaneOccupancy struct {

	// **参数说明**：车道编号。车道从左到右，从0开始编号
	LaneId *int32 `json:"lane_id,omitempty"`

	// **参数说明**：车道的空间占有率。
	SpaceOccupancy *float64 `json:"space_occupancy,omitempty"`

	// **参数说明**：车道的时间占有率。
	TimeOccupancy *float64 `json:"time_occupancy,omitempty"`
}

func (o LaneOccupancy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LaneOccupancy struct{}"
	}

	return strings.Join([]string{"LaneOccupancy", string(data)}, " ")
}
