package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListRelationResponse struct {
	// 关联工单列表

	CaseRelationList *[]CaseRealtionInfo `json:"case_relation_list,omitempty"`
	HttpStatusCode   int                 `json:"-"`
}

func (o ListRelationResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListRelationResponse struct{}"
	}

	return strings.Join([]string{"ListRelationResponse", string(data)}, " ")
}
