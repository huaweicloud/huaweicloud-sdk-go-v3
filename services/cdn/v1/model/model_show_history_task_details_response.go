package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowHistoryTaskDetailsResponse struct {
	// 任务id。

	Id *string `json:"id,omitempty"`
	// 刷新的类型， 其值可以为refresh或preheating。

	TaskType *string `json:"task_type,omitempty"`
	// 刷新结果。task_done表示刷新成功，task_inprocess表示刷新中。

	Status *string `json:"status,omitempty"`
	// 本次刷新时提交的URL列表。

	Urls *[]UrlObject `json:"urls,omitempty"`
	// 创建时间。

	CreateTime *int64 `json:"create_time,omitempty"`
	// 处理中的url个数。

	Processing *int32 `json:"processing,omitempty"`
	// 成功处理的url个数。

	Succeed *int32 `json:"succeed,omitempty"`
	// 处理失败的url个数。

	Failed *int32 `json:"failed,omitempty"`
	// 历史任务的url个数。

	Total *int32 `json:"total,omitempty"`
	// 默认是文件file,file：文件,directory：目录。

	FileType *string `json:"file_type,omitempty"`
	// 任务id

	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowHistoryTaskDetailsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowHistoryTaskDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowHistoryTaskDetailsResponse", string(data)}, " ")
}
