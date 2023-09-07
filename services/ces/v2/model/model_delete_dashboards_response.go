package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDashboardsResponse Response Object
type DeleteDashboardsResponse struct {

	// 批量删除监控看板返回结果
	Dashboards     *[]BatchDeleteDashboardRespInfo `json:"dashboards,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o DeleteDashboardsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDashboardsResponse struct{}"
	}

	return strings.Join([]string{"DeleteDashboardsResponse", string(data)}, " ")
}
