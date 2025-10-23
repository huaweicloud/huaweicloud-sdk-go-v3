package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReportTaskDetailsDataResponse Response Object
type ShowReportTaskDetailsDataResponse struct {

	// 本次分页查询响应的任务条数
	Count *int64 `json:"count,omitempty"`

	// 下一页的marker
	NextMarker *string `json:"next_marker,omitempty"`

	// 任务列表
	Tasks          *[]TaskEntity `json:"tasks,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowReportTaskDetailsDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReportTaskDetailsDataResponse struct{}"
	}

	return strings.Join([]string{"ShowReportTaskDetailsDataResponse", string(data)}, " ")
}
