package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteKillAllProcessResponse Response Object
type ExecuteKillAllProcessResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExecuteKillAllProcessResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteKillAllProcessResponse struct{}"
	}

	return strings.Join([]string{"ExecuteKillAllProcessResponse", string(data)}, " ")
}
