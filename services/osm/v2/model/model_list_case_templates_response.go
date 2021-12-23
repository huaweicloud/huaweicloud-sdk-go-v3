package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListCaseTemplatesResponse struct {
	// 总数

	TotalCount *int32 `json:"total_count,omitempty"`
	// 模板列表

	IncidentTemplateList *[]IncidentTempV2 `json:"incident_template_list,omitempty"`
	HttpStatusCode       int               `json:"-"`
}

func (o ListCaseTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCaseTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListCaseTemplatesResponse", string(data)}, " ")
}
