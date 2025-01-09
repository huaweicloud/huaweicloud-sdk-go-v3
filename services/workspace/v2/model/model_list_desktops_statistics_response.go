package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDesktopsStatisticsResponse Response Object
type ListDesktopsStatisticsResponse struct {

	// 桌面总数。
	TotalNum *int32 `json:"total_num,omitempty"`

	AttachStatistics *AttachStatistics `json:"attach_statistics,omitempty"`

	RunStateStatistics *DesktopRunStatisticsRsp `json:"run_state_statistics,omitempty"`

	LoginStateStatistics *DesktopLoginStatisticsRsp `json:"login_state_statistics,omitempty"`

	// 每个桌面池的情况统计，当desktop_type指定为POOL时返回
	DesktopPoolStatistics *[]DesktopPoolStatistics `json:"desktop_pool_statistics,omitempty"`
	HttpStatusCode        int                      `json:"-"`
}

func (o ListDesktopsStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopsStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListDesktopsStatisticsResponse", string(data)}, " ")
}
