package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchGaussMySqlProxyBinlogPullResponse Response Object
type SwitchGaussMySqlProxyBinlogPullResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchGaussMySqlProxyBinlogPullResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchGaussMySqlProxyBinlogPullResponse struct{}"
	}

	return strings.Join([]string{"SwitchGaussMySqlProxyBinlogPullResponse", string(data)}, " ")
}
