package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopJobRequestBody 停止构建任务接口请求体
type StopJobRequestBody struct {

	// 构建任务ID；编辑构建任务时，浏览器URL末尾的32位数字、字母组合的字符串
	JobId string `json:"job_id"`

	// 构建任务的构建编号，从1开始，每次构建递增1
	BuildNo string `json:"build_no"`
}

func (o StopJobRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopJobRequestBody struct{}"
	}

	return strings.Join([]string{"StopJobRequestBody", string(data)}, " ")
}
