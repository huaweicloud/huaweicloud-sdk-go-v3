package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMonitorSystemResponse Response Object
type CreateMonitorSystemResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateMonitorSystemResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMonitorSystemResponse struct{}"
	}

	return strings.Join([]string{"CreateMonitorSystemResponse", string(data)}, " ")
}
