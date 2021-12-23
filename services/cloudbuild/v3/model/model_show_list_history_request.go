package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowListHistoryRequest struct {
	// 构建的任务ID [获取项目下构建任务列表](https://support.huaweicloud.com/api-codeci/ShowJobListByProjectId.html)； 编辑构建任务时，浏览器URL末尾的32位数字、字母组合的字符串。

	JobId string `json:"job_id"`
	// 分页页码， 表示从此页开始查询， offset大于等于0

	Offset int32 `json:"offset"`
	// 每页显示的条目数量，limit小于等于100

	Limit int32 `json:"limit"`
	// 距今天的时间区间（单位：天），interval小于等于30

	Interval int32 `json:"interval"`
}

func (o ShowListHistoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowListHistoryRequest struct{}"
	}

	return strings.Join([]string{"ShowListHistoryRequest", string(data)}, " ")
}
