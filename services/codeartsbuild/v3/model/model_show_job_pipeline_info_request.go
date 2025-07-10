package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobPipelineInfoRequest Request Object
type ShowJobPipelineInfoRequest struct {

	// 构建的任务ID； 编辑构建任务时，浏览器URL末尾的32位数字、字母组合的字符串。
	JobId string `json:"job_id"`

	// 输入\"true\"或者\"false\"来控制返回参数是不是完整的
	All *string `json:"all,omitempty"`

	// 移除未使用的参数
	CheckParamUsed *string `json:"check_param_used,omitempty"`
}

func (o ShowJobPipelineInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobPipelineInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowJobPipelineInfoRequest", string(data)}, " ")
}
