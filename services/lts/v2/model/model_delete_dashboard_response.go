package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDashboardResponse Response Object
type DeleteDashboardResponse struct {

	// 删除仪表盘结果
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDashboardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDashboardResponse struct{}"
	}

	return strings.Join([]string{"DeleteDashboardResponse", string(data)}, " ")
}
