package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDiagnoseRecordsResponse struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 总条数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 获取的诊断记录列表
	Records        *[]DiagnoseRecordVo `json:"records,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListDiagnoseRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDiagnoseRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListDiagnoseRecordsResponse", string(data)}, " ")
}
