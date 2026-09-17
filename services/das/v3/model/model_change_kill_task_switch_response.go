package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeKillTaskSwitchResponse Response Object
type ChangeKillTaskSwitchResponse struct {

	// 操作是否成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ChangeKillTaskSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeKillTaskSwitchResponse struct{}"
	}

	return strings.Join([]string{"ChangeKillTaskSwitchResponse", string(data)}, " ")
}
