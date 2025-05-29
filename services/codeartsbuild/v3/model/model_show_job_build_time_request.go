package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobBuildTimeRequest Request Object
type ShowJobBuildTimeRequest struct {

	// 构建的任务ID； 编辑构建任务时，浏览器URL末尾的32位数字、字母组合的字符串。
	JobId string `json:"job_id"`

	// 代码仓名称。
	RepositoryName string `json:"repository_name"`

	// 代码仓分支名。
	Branch *string `json:"branch,omitempty"`

	// 查询时间，查最近几天的。
	Interval int32 `json:"interval"`
}

func (o ShowJobBuildTimeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobBuildTimeRequest struct{}"
	}

	return strings.Join([]string{"ShowJobBuildTimeRequest", string(data)}, " ")
}
