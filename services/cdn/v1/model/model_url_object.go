package model

import (
	"encoding/json"

	"strings"
)

type UrlObject struct {
	// 任务id

	Id *string `json:"id,omitempty"`
	// url的地址。

	Url *string `json:"url,omitempty"`
	// url的状态 processing， succeed， failed，分别表示处理中，完成，失败。

	Status *string `json:"status,omitempty"`
	// url创建时间，相对于UTC 1970-01-01到当前时间相隔的毫秒数。

	CreateTime *int64 `json:"create_time,omitempty"`
	// url所属task的id。

	TaskId *string `json:"task_id,omitempty"`
	// 标记处理原因。

	ProcessReason *string `json:"process_reason,omitempty"`
}

func (o UrlObject) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UrlObject struct{}"
	}

	return strings.Join([]string{"UrlObject", string(data)}, " ")
}
