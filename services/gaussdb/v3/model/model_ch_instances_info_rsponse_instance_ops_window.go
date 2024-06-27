package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChInstancesInfoRsponseInstanceOpsWindow 实例操作时间窗。
type ChInstancesInfoRsponseInstanceOpsWindow struct {

	// 时间窗周期。
	Period *string `json:"period,omitempty"`

	// 时间窗开始时间。
	StartTime string `json:"start_time"`

	// 时间窗结束时间。
	EndTime string `json:"end_time"`
}

func (o ChInstancesInfoRsponseInstanceOpsWindow) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChInstancesInfoRsponseInstanceOpsWindow struct{}"
	}

	return strings.Join([]string{"ChInstancesInfoRsponseInstanceOpsWindow", string(data)}, " ")
}
