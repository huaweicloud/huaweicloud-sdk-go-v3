package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRemoveTestCasesFromIteratorResponse Response Object
type BatchRemoveTestCasesFromIteratorResponse struct {

	// 操作的id, 由projectUuid + - + iteratorUri + - + caseId 组成
	Id *string `json:"id,omitempty"`

	// 操作名称
	Name           *string `json:"name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchRemoveTestCasesFromIteratorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRemoveTestCasesFromIteratorResponse struct{}"
	}

	return strings.Join([]string{"BatchRemoveTestCasesFromIteratorResponse", string(data)}, " ")
}
