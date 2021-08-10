package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowListPeriodHistoryRequest struct {
	// 构建的任务ID [获取项目下构建任务列表](https://support.huaweicloud.com/api-codeci/ShowJobListByProjectId.html)； 编辑构建任务时，浏览器URL末尾的32位数字、字母组合的字符串。

	JobId string `json:"job_id"`
	// 分页页码， 表示从此页开始查询， offset大于等于0

	Offset int32 `json:"offset"`
	// 每页显示的条目数量，limit小于等于100

	Limit int32 `json:"limit"`
	// 区间开始时间，格式yyyy-MM-dd。 开始时间和结束时间间隔不能超过30天

	StartTime string `json:"start_time"`
	// 区间结束时间，格式yyyy-MM-dd。 开始时间和结束时间间隔不能超过30天

	EndTime string `json:"end_time"`
}

func (o ShowListPeriodHistoryRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowListPeriodHistoryRequest struct{}"
	}

	return strings.Join([]string{"ShowListPeriodHistoryRequest", string(data)}, " ")
}
