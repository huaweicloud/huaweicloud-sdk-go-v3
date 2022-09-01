package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListCompareResultResponse struct {

	// 任务id。
	JobId *string `json:"job_id,omitempty" xml:"job_id"`

	ObjectLevelCompareResults *ObjectCompareResult `json:"object_level_compare_results,omitempty" xml:"object_level_compare_results"`

	LineCompareResults *LineCompareResult `json:"line_compare_results,omitempty" xml:"line_compare_results"`

	ContentCompareResults *ContentCompareResult `json:"content_compare_results,omitempty" xml:"content_compare_results"`

	CompareTaskListResults *CompareTaskListResult `json:"compare_task_list_results,omitempty" xml:"compare_task_list_results"`

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	// 错误信息。
	ErrorMsg       *string `json:"error_msg,omitempty" xml:"error_msg"`
	HttpStatusCode int     `json:"-"`
}

func (o ListCompareResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCompareResultResponse struct{}"
	}

	return strings.Join([]string{"ListCompareResultResponse", string(data)}, " ")
}
