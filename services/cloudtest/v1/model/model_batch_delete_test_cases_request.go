package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteTestCasesRequest Request Object
type BatchDeleteTestCasesRequest struct {

	// 是否异步
	IsAsync *bool `json:"is_async,omitempty"`

	Body *DeleteTestCaseInfo `json:"body,omitempty"`
}

func (o BatchDeleteTestCasesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteTestCasesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteTestCasesRequest", string(data)}, " ")
}
