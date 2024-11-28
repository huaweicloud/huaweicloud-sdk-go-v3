package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgentAddPathReq
type AgentAddPathReq struct {

	// 增加备份路径详情
	AddPath []string `json:"add_path"`

	// 增加排除目录 > 该特性目前处于公测阶段，部分region可能无法使用。
	ExcludePath *[]ExcludePath `json:"exclude_path,omitempty"`
}

func (o AgentAddPathReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentAddPathReq struct{}"
	}

	return strings.Join([]string{"AgentAddPathReq", string(data)}, " ")
}
