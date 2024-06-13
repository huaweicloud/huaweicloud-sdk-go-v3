package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRelationsByOneCaseRequest Request Object
type CreateRelationsByOneCaseRequest struct {

	// 用例uri
	CaseId string `json:"case_id"`

	Body *AddRelationsInfo `json:"body,omitempty"`
}

func (o CreateRelationsByOneCaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRelationsByOneCaseRequest struct{}"
	}

	return strings.Join([]string{"CreateRelationsByOneCaseRequest", string(data)}, " ")
}
