package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type KillProcessesOpenRequest struct {

	// 会话id集合
	ProcessIds []string `json:"process_ids"`
}

func (o KillProcessesOpenRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KillProcessesOpenRequest struct{}"
	}

	return strings.Join([]string{"KillProcessesOpenRequest", string(data)}, " ")
}
