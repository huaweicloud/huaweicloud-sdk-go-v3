package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgentRemovePathReq
type AgentRemovePathReq struct {

	// 移除备份路径详情
	RemovePath []string `json:"remove_path"`
}

func (o AgentRemovePathReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentRemovePathReq struct{}"
	}

	return strings.Join([]string{"AgentRemovePathReq", string(data)}, " ")
}
