package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRestoreRecordsResponse struct {
	// 实例恢复记录的详情数组。

	RestoreRecordResponse *[]InstanceRestoreInfo `json:"restore_record_response,omitempty"`
	// 返回记录数。

	TotalNum       *int32 `json:"total_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListRestoreRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRestoreRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListRestoreRecordsResponse", string(data)}, " ")
}
