package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLogicalProcessesResponse Response Object
type ShowLogicalProcessesResponse struct {

	// 逻辑会话列表
	LogicalProcesses *[]LogicalProcessInfo `json:"logical_processes,omitempty"`

	// 总数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowLogicalProcessesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLogicalProcessesResponse struct{}"
	}

	return strings.Join([]string{"ShowLogicalProcessesResponse", string(data)}, " ")
}
