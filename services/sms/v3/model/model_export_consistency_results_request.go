package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportConsistencyResultsRequest Request Object
type ExportConsistencyResultsRequest struct {

	// 中英文选择。当前仅支持英文。
	XLanguage *string `json:"X-Language,omitempty"`

	Body *BatchGetConsistencyResultReq `json:"body,omitempty"`
}

func (o ExportConsistencyResultsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportConsistencyResultsRequest struct{}"
	}

	return strings.Join([]string{"ExportConsistencyResultsRequest", string(data)}, " ")
}
