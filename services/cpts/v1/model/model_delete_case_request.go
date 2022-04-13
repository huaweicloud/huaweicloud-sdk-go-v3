package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteCaseRequest struct {
	// 用例id

	CaseId int32 `json:"case_id"`
}

func (o DeleteCaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCaseRequest struct{}"
	}

	return strings.Join([]string{"DeleteCaseRequest", string(data)}, " ")
}
