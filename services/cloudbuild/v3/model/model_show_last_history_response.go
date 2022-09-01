package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowLastHistoryResponse struct {

	// 构建记录id--唯一key
	RecordId *string `json:"record_id,omitempty" xml:"record_id"`

	// 构建任务ID
	JobId *string `json:"job_id,omitempty" xml:"job_id"`

	// 构建任务名称
	JobName *string `json:"job_name,omitempty" xml:"job_name"`

	// 构建编号
	BuildNumber *int32 `json:"build_number,omitempty" xml:"build_number"`

	// 构建开始时间
	StartTime *string `json:"start_time,omitempty" xml:"start_time"`

	// 构建结束时间
	EndTime *string `json:"end_time,omitempty" xml:"end_time"`

	// 构建执行结果
	Result *string `json:"result,omitempty" xml:"result"`

	// commitId
	CommitId       *string `json:"commit_id,omitempty" xml:"commit_id"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowLastHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLastHistoryResponse struct{}"
	}

	return strings.Join([]string{"ShowLastHistoryResponse", string(data)}, " ")
}
