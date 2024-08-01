package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteKillLogicalProcessesResponse Response Object
type ExecuteKillLogicalProcessesResponse struct {

	// 操作结果。
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteKillLogicalProcessesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteKillLogicalProcessesResponse struct{}"
	}

	return strings.Join([]string{"ExecuteKillLogicalProcessesResponse", string(data)}, " ")
}
