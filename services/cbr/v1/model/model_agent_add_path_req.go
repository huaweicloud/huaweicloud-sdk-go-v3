package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type AgentAddPathReq struct {

	// 增加备份路径详情
	AddPath []string `json:"add_path"`
}

func (o AgentAddPathReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentAddPathReq struct{}"
	}

	return strings.Join([]string{"AgentAddPathReq", string(data)}, " ")
}
