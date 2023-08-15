package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMonitorLogResponse Response Object
type ListMonitorLogResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 当前页日志数量
	Size *int32 `json:"size,omitempty"`

	// 任务监控日志当前页元素
	Entities       *[]TaskMonitorLog `json:"entities,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListMonitorLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMonitorLogResponse struct{}"
	}

	return strings.Join([]string{"ListMonitorLogResponse", string(data)}, " ")
}
