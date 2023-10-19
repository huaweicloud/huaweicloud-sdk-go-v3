package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResumeBuildJobRequest Request Object
type ResumeBuildJobRequest struct {

	// 构建的任务ID [获取项目下构建任务列表](https://support.huaweicloud.com/api-codeci/ShowJobListByProjectId.html)； 编辑构建任务时，浏览器URL末尾的32位数字、字母组合的字符串。
	JobId string `json:"job_id"`
}

func (o ResumeBuildJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResumeBuildJobRequest struct{}"
	}

	return strings.Join([]string{"ResumeBuildJobRequest", string(data)}, " ")
}
