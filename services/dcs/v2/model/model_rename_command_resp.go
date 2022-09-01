package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RenameCommandResp struct {

	// 命令command
	Command *string `json:"command,omitempty" xml:"command"`

	// 命令flushall
	Flushall *string `json:"flushall,omitempty" xml:"flushall"`

	// 命令flushdb
	Flushdb *string `json:"flushdb,omitempty" xml:"flushdb"`

	// 命令hgetall
	Hgetall *string `json:"hgetall,omitempty" xml:"hgetall"`

	// 命令keys
	Keys *string `json:"keys,omitempty" xml:"keys"`
}

func (o RenameCommandResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RenameCommandResp struct{}"
	}

	return strings.Join([]string{"RenameCommandResp", string(data)}, " ")
}
