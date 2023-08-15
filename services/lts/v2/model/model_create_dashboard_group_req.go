package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateDashboardGroupReq struct {

	// 仪表盘分组名称
	GroupName string `json:"group_name"`
}

func (o CreateDashboardGroupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDashboardGroupReq struct{}"
	}

	return strings.Join([]string{"CreateDashboardGroupReq", string(data)}, " ")
}
