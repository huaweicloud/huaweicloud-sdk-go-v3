package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DashBoardIdItem struct {

	// 监控看板id
	DashboardId *string `json:"dashboard_id,omitempty"`
}

func (o DashBoardIdItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DashBoardIdItem struct{}"
	}

	return strings.Join([]string{"DashBoardIdItem", string(data)}, " ")
}
