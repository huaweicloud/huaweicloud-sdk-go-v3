package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobsRequest Request Object
type ListJobsRequest struct {

	// **参数解释**：每页返回的个数。  **约束限制**：不涉及  **取值范围**：0-2000  **默认取值**：2000
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：上一页最后一条记录的ID。  **约束限制**： - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。  **取值范围**：不涉及  **默认取值**：不涉及
	Marker *string `json:"marker,omitempty"`

	// **参数解释**：是否反向查询。  **约束限制**： - 必须与limit一起使用。 - 当page_reverse=true时，若要查询上一页，marker取值为当前页返回值的previous_marker。  **取值范围**： - true：查询上一页。 - false：查询下一页。  **默认取值**：false
	PageReverse *bool `json:"page_reverse,omitempty"`

	// **参数解释**：任务ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**：任务类型。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	JobType *string `json:"job_type,omitempty"`

	// **参数解释**：任务状态。  **约束限制**：不涉及  **取值范围**：INIT,RUNNING,FAIL,SUCCESS,ROLLBACKING,COMPLETE,ROLLBACK_FAIL,CANCEL  **默认取值**：不涉及
	Status *string `json:"status,omitempty"`

	// **参数解释**：任务的错误码。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	ErrorCode *string `json:"error_code,omitempty"`

	// **参数解释**：资源ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	ResourceId *string `json:"resource_id,omitempty"`

	// **参数解释**：查询任务的开始时间大于等于传入时间的任务。格式：yyyy-MM-dd'T'HH:mm:ss  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	BeginTime *string `json:"begin_time,omitempty"`
}

func (o ListJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobsRequest struct{}"
	}

	return strings.Join([]string{"ListJobsRequest", string(data)}, " ")
}
