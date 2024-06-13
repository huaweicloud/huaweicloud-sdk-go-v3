package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRelationsByOneCaseRequest Request Object
type DeleteRelationsByOneCaseRequest struct {

	// 用例id
	CaseId string `json:"case_id"`

	Body *DeleteRelationsByOneCaseInfo `json:"body,omitempty"`
}

func (o DeleteRelationsByOneCaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRelationsByOneCaseRequest struct{}"
	}

	return strings.Join([]string{"DeleteRelationsByOneCaseRequest", string(data)}, " ")
}
