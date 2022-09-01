package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListMonitorInfosResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty" xml:"total"`

	// 当前页监控信息数量
	Size *int32 `json:"size,omitempty" xml:"size"`

	// 任务监控信息当前页元素
	Entities       *[]TaskMonitorInfo `json:"entities,omitempty" xml:"entities"`
	HttpStatusCode int                `json:"-"`
}

func (o ListMonitorInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMonitorInfosResponse struct{}"
	}

	return strings.Join([]string{"ListMonitorInfosResponse", string(data)}, " ")
}
