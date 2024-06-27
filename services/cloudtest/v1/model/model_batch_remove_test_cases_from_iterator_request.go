package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRemoveTestCasesFromIteratorRequest Request Object
type BatchRemoveTestCasesFromIteratorRequest struct {

	// 迭代uri
	IteratorId string `json:"iterator_id"`

	Body *CaseRemoveInfo `json:"body,omitempty"`
}

func (o BatchRemoveTestCasesFromIteratorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRemoveTestCasesFromIteratorRequest struct{}"
	}

	return strings.Join([]string{"BatchRemoveTestCasesFromIteratorRequest", string(data)}, " ")
}
