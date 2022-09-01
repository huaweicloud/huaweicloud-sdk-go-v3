package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateCompareTaskResponse struct {

	// 任务id。
	JobId *string `json:"job_id,omitempty" xml:"job_id"`

	ObjectLevelCompareCreateResult *CreateCompareTaskResult `json:"object_level_compare_create_result,omitempty" xml:"object_level_compare_create_result"`

	DataLevelCompareCreateResult *CreateCompareTaskResult `json:"data_level_compare_create_result,omitempty" xml:"data_level_compare_create_result"`

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	// 错误信息。
	ErrorMsg       *string `json:"error_msg,omitempty" xml:"error_msg"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCompareTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCompareTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateCompareTaskResponse", string(data)}, " ")
}
