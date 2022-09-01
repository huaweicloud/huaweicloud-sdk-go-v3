package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowUrlTaskInfoRequest struct {

	// 起始时间戳（毫秒），默认当天00:00
	StartTime *int32 `json:"start_time,omitempty" xml:"start_time"`

	// 结束时间戳（毫秒），默认次日00:00
	EndTime *int32 `json:"end_time,omitempty" xml:"end_time"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 单次查询数据条数，上限为100
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 刷新预热url
	Url *string `json:"url,omitempty" xml:"url"`

	// 任务类型，REFRESH：刷新任务；PREHEATING：预热任务
	TaskType *string `json:"task_type,omitempty" xml:"task_type"`

	// url状态，状态类型：processing：处理中；succeed：完成；failed：失败；waiting：等待；refreshing：刷新中; preheating : 预热中
	Status *string `json:"status,omitempty" xml:"status"`

	// 文件类型，file:文件;directory:目录
	FileType *string `json:"file_type,omitempty" xml:"file_type"`
}

func (o ShowUrlTaskInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUrlTaskInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowUrlTaskInfoRequest", string(data)}, " ")
}
