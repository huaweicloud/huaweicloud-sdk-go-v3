package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveTemplateUsedInfoRequestBody 保存模板使用记录接口请求体
type SaveTemplateUsedInfoRequestBody struct {

	// 构建任务ID；编辑构建任务时，浏览器URL末尾的32位数字、字母组合的字符串
	JobId string `json:"job_id"`

	// 构建模版ID
	TemplateId string `json:"template_id"`
}

func (o SaveTemplateUsedInfoRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveTemplateUsedInfoRequestBody struct{}"
	}

	return strings.Join([]string{"SaveTemplateUsedInfoRequestBody", string(data)}, " ")
}
