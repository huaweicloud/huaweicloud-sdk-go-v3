package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOsStatisticsInfoResponse Response Object
type ShowOsStatisticsInfoResponse struct {

	// win_num
	WinNum *int32 `json:"win_num,omitempty"`

	// linux_num
	LinuxNum *int32 `json:"linux_num,omitempty"`

	// 操作系统统计数据列表
	OsList         *[]OsStatisticsInfo `json:"os_list,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowOsStatisticsInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOsStatisticsInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowOsStatisticsInfoResponse", string(data)}, " ")
}
