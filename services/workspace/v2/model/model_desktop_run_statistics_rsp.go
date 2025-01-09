package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DesktopRunStatisticsRsp 桌面运行状态。
type DesktopRunStatisticsRsp struct {

	// 停止个数。
	StopNum int32 `json:"stop_num"`

	// 运行中个数。
	ActiveNum int32 `json:"active_num"`

	// 故障个数。
	ErrorNum int32 `json:"error_num"`

	// 休眠个数。
	HibernatedNum int32 `json:"hibernated_num"`
}

func (o DesktopRunStatisticsRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DesktopRunStatisticsRsp struct{}"
	}

	return strings.Join([]string{"DesktopRunStatisticsRsp", string(data)}, " ")
}
