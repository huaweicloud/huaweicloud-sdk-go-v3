package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DesktopVersionStatistic 单个版本的桌面统计信息。
type DesktopVersionStatistic struct {

	// 桌面版本号。
	Version *string `json:"version,omitempty"`

	// 操作系统。
	OsType *string `json:"os_type,omitempty"`

	// 该版本下的桌面数量。
	DesktopCount *int32 `json:"desktop_count,omitempty"`
}

func (o DesktopVersionStatistic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DesktopVersionStatistic struct{}"
	}

	return strings.Join([]string{"DesktopVersionStatistic", string(data)}, " ")
}
