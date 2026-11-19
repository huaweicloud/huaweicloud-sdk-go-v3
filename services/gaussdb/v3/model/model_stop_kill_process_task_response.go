package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopKillProcessTaskResponse Response Object
type StopKillProcessTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopKillProcessTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopKillProcessTaskResponse struct{}"
	}

	return strings.Join([]string{"StopKillProcessTaskResponse", string(data)}, " ")
}
