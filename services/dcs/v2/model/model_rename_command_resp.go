package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RenameCommandResp struct {

	// 命令command
	Command *string `json:"command,omitempty"`

	// 命令flushall
	Flushall *string `json:"flushall,omitempty"`

	// 命令flushdb
	Flushdb *string `json:"flushdb,omitempty"`

	// 命令hgetall
	Hgetall *string `json:"hgetall,omitempty"`

	// 命令keys
	Keys *string `json:"keys,omitempty"`

	// 命令hscan
	Hscan *string `json:"hscan,omitempty"`

	// 命令scan
	Scan *string `json:"scan,omitempty"`

	// 命令sscan
	Sscan *string `json:"sscan,omitempty"`

	// 命令zscan
	Zscan *string `json:"zscan,omitempty"`
}

func (o RenameCommandResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RenameCommandResp struct{}"
	}

	return strings.Join([]string{"RenameCommandResp", string(data)}, " ")
}
