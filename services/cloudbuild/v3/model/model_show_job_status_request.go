package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowJobStatusRequest struct {
	// 构建的任务ID [获取项目下构建任务列表](https://support.huaweicloud.com/api-codeci/ShowJobListByProjectId.html)； 编辑构建任务时，浏览器URL末尾的32位数字、字母组合的字符串。

	JobId string `json:"job_id"`
}

func (o ShowJobStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowJobStatusRequest", string(data)}, " ")
}
