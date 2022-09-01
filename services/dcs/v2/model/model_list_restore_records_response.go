package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRestoreRecordsResponse struct {

	// 实例恢复记录的详情数组。
	RestoreRecordResponse *[]InstanceRestoreInfo `json:"restore_record_response,omitempty" xml:"restore_record_response"`

	// 返回记录数。
	TotalNum       *int32 `json:"total_num,omitempty" xml:"total_num"`
	HttpStatusCode int    `json:"-"`
}

func (o ListRestoreRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRestoreRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListRestoreRecordsResponse", string(data)}, " ")
}
