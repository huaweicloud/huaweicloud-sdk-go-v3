package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WorkitemStatusRecords struct {

	// 工作项的记录id，一个工作项对应一条记录
	WorkItemRecordId *string `json:"work_item_record_id,omitempty"`

	// 工作项id
	WorkItemId *string `json:"work_item_id,omitempty"`

	// devcloud项目的32位id
	ProjectId *string `json:"project_id,omitempty"`

	// 操作历史
	WorkItemStatuses *[]WorkitemStatus `json:"work_item_statuses,omitempty"`
}

func (o WorkitemStatusRecords) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkitemStatusRecords struct{}"
	}

	return strings.Join([]string{"WorkitemStatusRecords", string(data)}, " ")
}
