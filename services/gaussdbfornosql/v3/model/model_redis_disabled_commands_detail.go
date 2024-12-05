package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RedisDisabledCommandsDetail struct {

	// key所在的DB。
	DbId string `json:"db_id"`

	// key名。
	Key string `json:"key"`

	// 命令列表。
	Commands []string `json:"commands"`
}

func (o RedisDisabledCommandsDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedisDisabledCommandsDetail struct{}"
	}

	return strings.Join([]string{"RedisDisabledCommandsDetail", string(data)}, " ")
}
