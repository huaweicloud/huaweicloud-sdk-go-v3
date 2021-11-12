package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 执行任务接口请求体
type RunJobRequestBody struct {
	// 构建任务ID；编辑构建任务时，浏览器URL末尾的32位数字、字母组合的字符串

	JobId string `json:"job_id"`
	// 自定义参数

	Parameter *[]ParameterItem `json:"parameter,omitempty"`

	Scm *Scm `json:"scm,omitempty"`
}

func (o RunJobRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunJobRequestBody struct{}"
	}

	return strings.Join([]string{"RunJobRequestBody", string(data)}, " ")
}
