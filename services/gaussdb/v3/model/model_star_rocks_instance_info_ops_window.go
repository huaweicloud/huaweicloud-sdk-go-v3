package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StarRocksInstanceInfoOpsWindow 实例操作时间窗。
type StarRocksInstanceInfoOpsWindow struct {

	// 时间窗周期。
	Period *string `json:"period,omitempty"`

	// 时间窗开始时间。
	StartTime *string `json:"start_time,omitempty"`

	// 时间窗结束时间。
	EndTime *string `json:"end_time,omitempty"`
}

func (o StarRocksInstanceInfoOpsWindow) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StarRocksInstanceInfoOpsWindow struct{}"
	}

	return strings.Join([]string{"StarRocksInstanceInfoOpsWindow", string(data)}, " ")
}
