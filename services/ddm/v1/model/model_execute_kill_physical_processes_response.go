package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteKillPhysicalProcessesResponse Response Object
type ExecuteKillPhysicalProcessesResponse struct {

	// 操作结果。
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteKillPhysicalProcessesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteKillPhysicalProcessesResponse struct{}"
	}

	return strings.Join([]string{"ExecuteKillPhysicalProcessesResponse", string(data)}, " ")
}
