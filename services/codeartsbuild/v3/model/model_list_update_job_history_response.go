package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUpdateJobHistoryResponse Response Object
type ListUpdateJobHistoryResponse struct {

	// 状态
	Status *string `json:"status,omitempty"`

	// 错误
	Error *interface{} `json:"error,omitempty"`

	Result         *JobUpdateRecordListVoResult `json:"result,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListUpdateJobHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUpdateJobHistoryResponse struct{}"
	}

	return strings.Join([]string{"ListUpdateJobHistoryResponse", string(data)}, " ")
}
