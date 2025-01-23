package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobsRequest Request Object
type ListJobsRequest struct {

	// 参数解释：任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 参数解释：任务类型。
	JobType *string `json:"job_type,omitempty"`

	// 参数解释：任务状态。  取值范围：INIT,RUNNING,FAIL,SUCCESS,ROLLBACKING,COMPLETE,ROLLBACK_FAIL,CANCEL
	Status *string `json:"status,omitempty"`

	// 参数解释： 任务的错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 参数解释：资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 参数解释：查询任务的开始时间大于等于传入时间的任务。格式：yyyy-MM-dd'T'HH:mm:ss
	BeginTime *string `json:"begin_time,omitempty"`
}

func (o ListJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobsRequest struct{}"
	}

	return strings.Join([]string{"ListJobsRequest", string(data)}, " ")
}
