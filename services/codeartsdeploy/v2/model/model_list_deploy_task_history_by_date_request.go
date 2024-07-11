package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDeployTaskHistoryByDateRequest Request Object
type ListDeployTaskHistoryByDateRequest struct {

	// 项目id
	ProjectId string `json:"project_id"`

	// 任务id
	Id string `json:"id"`

	// 分页页码， 表示从此页开始查询， page大于等于1
	Page int32 `json:"page"`

	// 每页显示的条目数量，size小于等于100
	Size int32 `json:"size"`

	// 区间开始时间，格式yyyy-MM-dd。 开始时间和结束时间间隔不能超过30天
	StartDate string `json:"start_date"`

	// 区间结束时间，格式yyyy-MM-dd。 开始时间和结束时间间隔不能超过30天
	EndDate string `json:"end_date"`
}

func (o ListDeployTaskHistoryByDateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDeployTaskHistoryByDateRequest struct{}"
	}

	return strings.Join([]string{"ListDeployTaskHistoryByDateRequest", string(data)}, " ")
}
