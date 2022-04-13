package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListCaseLabelsResponse struct {
	// 工单关联的标签列表

	CaseLabelList  *[]CaseLabelInfo `json:"case_label_list,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListCaseLabelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCaseLabelsResponse struct{}"
	}

	return strings.Join([]string{"ListCaseLabelsResponse", string(data)}, " ")
}
