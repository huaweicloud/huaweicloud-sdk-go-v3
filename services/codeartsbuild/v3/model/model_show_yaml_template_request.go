package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowYamlTemplateRequest Request Object
type ShowYamlTemplateRequest struct {

	// 构建的任务ID； 编辑构建任务时，浏览器URL末尾的32位数字、字母组合的字符串。
	JobId string `json:"job_id"`

	// 默认主机类型
	DefaultHost *string `json:"default_host,omitempty"`
}

func (o ShowYamlTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowYamlTemplateRequest struct{}"
	}

	return strings.Join([]string{"ShowYamlTemplateRequest", string(data)}, " ")
}
