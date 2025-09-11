package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DashBoardNameItemResp struct {

	// **参数解释** 自定义监控看板名称 **取值范围** 长度为[1,128]个字符，只允许中文、英文、数字0-9、_和-
	DashboardName *string `json:"dashboard_name,omitempty"`
}

func (o DashBoardNameItemResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DashBoardNameItemResp struct{}"
	}

	return strings.Join([]string{"DashBoardNameItemResp", string(data)}, " ")
}
