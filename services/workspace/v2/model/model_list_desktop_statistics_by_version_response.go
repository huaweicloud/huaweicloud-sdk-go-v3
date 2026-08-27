package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDesktopStatisticsByVersionResponse Response Object
type ListDesktopStatisticsByVersionResponse struct {

	// 按版本分组的桌面统计信息列表。
	VersionStatistics *[]DesktopVersionStatistic `json:"version_statistics,omitempty"`
	HttpStatusCode    int                        `json:"-"`
}

func (o ListDesktopStatisticsByVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopStatisticsByVersionResponse struct{}"
	}

	return strings.Join([]string{"ListDesktopStatisticsByVersionResponse", string(data)}, " ")
}
