package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveRedisDisabledCommandsResponse Response Object
type SaveRedisDisabledCommandsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SaveRedisDisabledCommandsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveRedisDisabledCommandsResponse struct{}"
	}

	return strings.Join([]string{"SaveRedisDisabledCommandsResponse", string(data)}, " ")
}
