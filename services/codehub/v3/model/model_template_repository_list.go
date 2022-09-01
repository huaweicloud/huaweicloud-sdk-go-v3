package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TemplateRepositoryList struct {

	// 模板列表
	Projects *[]TemplateRepository `json:"projects,omitempty" xml:"projects"`

	// 模板总数
	Total *int32 `json:"total,omitempty" xml:"total"`
}

func (o TemplateRepositoryList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateRepositoryList struct{}"
	}

	return strings.Join([]string{"TemplateRepositoryList", string(data)}, " ")
}
