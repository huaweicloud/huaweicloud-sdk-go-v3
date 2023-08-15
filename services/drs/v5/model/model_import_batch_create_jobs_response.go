package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportBatchCreateJobsResponse Response Object
type ImportBatchCreateJobsResponse struct {

	// 批量导入任务id。
	AsyncJobId *string `json:"async_job_id,omitempty"`

	// 导入失败的错误信息。
	ImportErrorMessages *[]ImportErrorMessageResp `json:"import_error_messages,omitempty"`

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述。
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ImportBatchCreateJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportBatchCreateJobsResponse struct{}"
	}

	return strings.Join([]string{"ImportBatchCreateJobsResponse", string(data)}, " ")
}
