package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDisasterRecoveryRecordResponse Response Object
type ListDisasterRecoveryRecordResponse struct {

	// **参数解释**: 操作记录。
	Records        *[]RecordInfoResponseResult `json:"records,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListDisasterRecoveryRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDisasterRecoveryRecordResponse struct{}"
	}

	return strings.Join([]string{"ListDisasterRecoveryRecordResponse", string(data)}, " ")
}
