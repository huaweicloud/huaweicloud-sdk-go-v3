package model

import (
	"encoding/json"

	"strings"
)

type UrlObject struct {
	// url的id

	Id *string `json:"id,omitempty"`
	// url的地址。

	Url *string `json:"url,omitempty"`
	// url的状态 processing 处理中，succeed 完成，failed 失败，waiting 等待，refreshing 刷新中，preheating 预热中。

	Status *string `json:"status,omitempty"`
	// url创建时间，相对于UTC 1970-01-01到当前时间相隔的毫秒数。

	CreateTime *int64 `json:"create_time,omitempty"`
	// 任务id。

	TaskId *string `json:"task_id,omitempty"`
}

func (o UrlObject) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UrlObject struct{}"
	}

	return strings.Join([]string{"UrlObject", string(data)}, " ")
}
