package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobConfigRequest Request Object
type ListJobConfigRequest struct {

	// 构建的任务ID； 编辑构建任务时，浏览器URL末尾的32位数字、字母组合的字符串。
	JobId string `json:"job_id"`

	// 输入\"true\"或者\"false\"来控制返回参数是不是完整的
	GetAllParams *string `json:"get_all_params,omitempty"`
}

func (o ListJobConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobConfigRequest struct{}"
	}

	return strings.Join([]string{"ListJobConfigRequest", string(data)}, " ")
}
