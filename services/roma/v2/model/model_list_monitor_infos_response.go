package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMonitorInfosResponse Response Object
type ListMonitorInfosResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 当前页监控信息数量
	Size *int32 `json:"size,omitempty"`

	// 任务监控信息当前页元素
	Entities       *[]TaskMonitorInfo `json:"entities,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListMonitorInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMonitorInfosResponse struct{}"
	}

	return strings.Join([]string{"ListMonitorInfosResponse", string(data)}, " ")
}
