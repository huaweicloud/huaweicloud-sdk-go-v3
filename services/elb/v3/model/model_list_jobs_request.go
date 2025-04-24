package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobsRequest Request Object
type ListJobsRequest struct {

	// 参数解释：每页返回的个数。  取值范围：0-2000  默认取值：2000
	Limit *int32 `json:"limit,omitempty"`

	// 上一页最后一条记录的ID。  使用说明： - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。
	Marker *string `json:"marker,omitempty"`

	// 是否反向查询。  取值： - true：查询上一页。 - false：查询下一页，默认。  使用说明： - 必须与limit一起使用。 - 当page_reverse=true时，若要查询上一页，marker取值为当前页返回值的previous_marker
	PageReverse *bool `json:"page_reverse,omitempty"`

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
