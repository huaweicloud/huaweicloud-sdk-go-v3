package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopKillTaskResponse Response Object
type StopKillTaskResponse struct {

	// 操作是否成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o StopKillTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopKillTaskResponse struct{}"
	}

	return strings.Join([]string{"StopKillTaskResponse", string(data)}, " ")
}
