package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTemplate 模板中构建执行步骤列表
type CreateTemplate struct {

	// 构建执行的步骤
	Steps []CreateTemplateSteps `json:"steps"`
}

func (o CreateTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTemplate struct{}"
	}

	return strings.Join([]string{"CreateTemplate", string(data)}, " ")
}
