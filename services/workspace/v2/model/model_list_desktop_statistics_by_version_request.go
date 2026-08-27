package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDesktopStatisticsByVersionRequest Request Object
type ListDesktopStatisticsByVersionRequest struct {

	// 桌面版本号列表，用于过滤指定版本的桌面统计。
	Versions *[]string `json:"versions,omitempty"`
}

func (o ListDesktopStatisticsByVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopStatisticsByVersionRequest struct{}"
	}

	return strings.Join([]string{"ListDesktopStatisticsByVersionRequest", string(data)}, " ")
}
