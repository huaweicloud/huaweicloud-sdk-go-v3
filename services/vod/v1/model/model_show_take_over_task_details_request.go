package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowTakeOverTaskDetailsRequest struct {
	// 起始时间.指定task_id时该参数无效<br/>

	TaskId *string `json:"task_id,omitempty"`
	// 分页编号,默认为0。<br/>

	Page *int32 `json:"page,omitempty"`
	// 每页记录数。默认10，范围[1,100]<br/>

	Size *int32 `json:"size,omitempty"`
}

func (o ShowTakeOverTaskDetailsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowTakeOverTaskDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowTakeOverTaskDetailsRequest", string(data)}, " ")
}
