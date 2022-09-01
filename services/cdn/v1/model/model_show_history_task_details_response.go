package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowHistoryTaskDetailsResponse struct {

	// 任务id。
	Id *string `json:"id,omitempty" xml:"id"`

	// 任务类型，REFRESH：刷新任务；PREHEATING：预热任务。
	TaskType *string `json:"task_type,omitempty" xml:"task_type"`

	// 任务执行结果,task_done:成功，task_inprocess:处理中。
	Status *string `json:"status,omitempty" xml:"status"`

	// 本次提交的url列表。
	Urls *[]UrlObject `json:"urls,omitempty" xml:"urls"`

	// 创建时间。
	CreateTime *int64 `json:"create_time,omitempty" xml:"create_time"`

	// 处理中的url个数。
	Processing *int32 `json:"processing,omitempty" xml:"processing"`

	// 成功处理的url个数。
	Succeed *int32 `json:"succeed,omitempty" xml:"succeed"`

	// 处理失败的url个数。
	Failed *int32 `json:"failed,omitempty" xml:"failed"`

	// 历史任务的url个数。
	Total *int32 `json:"total,omitempty" xml:"total"`

	// 文件类型，file：文件；directory：目录，默认是文件file,
	FileType       *string `json:"file_type,omitempty" xml:"file_type"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowHistoryTaskDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHistoryTaskDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowHistoryTaskDetailsResponse", string(data)}, " ")
}
