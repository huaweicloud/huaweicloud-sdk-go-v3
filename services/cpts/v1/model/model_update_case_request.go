package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCaseRequest Request Object
type UpdateCaseRequest struct {

	// 用例id
	CaseId int32 `json:"case_id"`

	// 类型
	Target string `json:"target"`

	Body *CaseInfoDetail `json:"body,omitempty"`
}

func (o UpdateCaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCaseRequest struct{}"
	}

	return strings.Join([]string{"UpdateCaseRequest", string(data)}, " ")
}
