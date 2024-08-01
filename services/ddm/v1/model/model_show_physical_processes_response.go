package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPhysicalProcessesResponse Response Object
type ShowPhysicalProcessesResponse struct {

	// 物理会话信息列表
	PhysicalProcesses *[]PhysicalProcessInfo `json:"physical_processes,omitempty"`

	// 总数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowPhysicalProcessesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPhysicalProcessesResponse struct{}"
	}

	return strings.Join([]string{"ShowPhysicalProcessesResponse", string(data)}, " ")
}
