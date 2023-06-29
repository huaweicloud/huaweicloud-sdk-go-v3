package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchToSlaveResponse Response Object
type SwitchToSlaveResponse struct {

	// 容灾实例主备倒换的工作ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchToSlaveResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchToSlaveResponse struct{}"
	}

	return strings.Join([]string{"SwitchToSlaveResponse", string(data)}, " ")
}
