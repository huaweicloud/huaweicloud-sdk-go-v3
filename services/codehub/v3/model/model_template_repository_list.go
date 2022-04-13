package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TemplateRepositoryList struct {
	// 模板列表

	Projects *[]TemplateRepository `json:"projects,omitempty"`
	// 模板总数

	Total *int32 `json:"total,omitempty"`
}

func (o TemplateRepositoryList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateRepositoryList struct{}"
	}

	return strings.Join([]string{"TemplateRepositoryList", string(data)}, " ")
}
