package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListMonitorLogResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty" xml:"total"`

	// 当前页日志数量
	Size *int32 `json:"size,omitempty" xml:"size"`

	// 任务监控日志当前页元素
	Entities       *[]TaskMonitorLog `json:"entities,omitempty" xml:"entities"`
	HttpStatusCode int               `json:"-"`
}

func (o ListMonitorLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMonitorLogResponse struct{}"
	}

	return strings.Join([]string{"ListMonitorLogResponse", string(data)}, " ")
}
