package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DashBoardNameItem struct {

	// 自定义监控看板名称
	DashboardName *string `json:"dashboard_name,omitempty"`
}

func (o DashBoardNameItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DashBoardNameItem struct{}"
	}

	return strings.Join([]string{"DashBoardNameItem", string(data)}, " ")
}
