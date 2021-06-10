package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListCaseLabelsResponse struct {
	// 工单关联的标签列表

	CaseLabelList  *[]CaseLabelInfo `json:"case_label_list,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListCaseLabelsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListCaseLabelsResponse struct{}"
	}

	return strings.Join([]string{"ListCaseLabelsResponse", string(data)}, " ")
}
