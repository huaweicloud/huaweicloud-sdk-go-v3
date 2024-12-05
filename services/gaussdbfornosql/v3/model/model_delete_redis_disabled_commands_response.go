package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRedisDisabledCommandsResponse Response Object
type DeleteRedisDisabledCommandsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteRedisDisabledCommandsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRedisDisabledCommandsResponse struct{}"
	}

	return strings.Join([]string{"DeleteRedisDisabledCommandsResponse", string(data)}, " ")
}
