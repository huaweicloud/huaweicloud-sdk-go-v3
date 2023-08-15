package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobSuccessRatioRequest Request Object
type ShowJobSuccessRatioRequest struct {

	// 构建的任务ID [获取项目下构建任务列表](https://support.huaweicloud.com/api-codeci/ShowJobListByProjectId.html)； 编辑构建任务时，浏览器URL末尾的32位数字、字母组合的字符串。
	JobId string `json:"job_id"`

	// 区间开始时间，格式yyyy-MM-dd。
	StartTime string `json:"start_time"`

	// 区间结束时间，格式yyyy-MM-dd。
	EndTime string `json:"end_time"`
}

func (o ShowJobSuccessRatioRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobSuccessRatioRequest struct{}"
	}

	return strings.Join([]string{"ShowJobSuccessRatioRequest", string(data)}, " ")
}
