package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DesktopPoolStatistics 桌面池状态统计。
type DesktopPoolStatistics struct {

	// 桌面池id。
	DesktopPoolId *string `json:"desktop_pool_id,omitempty"`

	// 桌面池名称。
	DesktopPoolName *string `json:"desktop_pool_name,omitempty"`

	// 桌面总数。
	TotalNum *int32 `json:"total_num,omitempty"`

	AttachStatistics *AttachStatistics `json:"attach_statistics,omitempty"`

	RunStateStatistics *DesktopRunStatisticsRsp `json:"run_state_statistics,omitempty"`

	LoginStateStatistics *DesktopLoginStatisticsRsp `json:"login_state_statistics,omitempty"`
}

func (o DesktopPoolStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DesktopPoolStatistics struct{}"
	}

	return strings.Join([]string{"DesktopPoolStatistics", string(data)}, " ")
}
