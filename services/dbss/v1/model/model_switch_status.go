package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SwitchStatus struct {

	// 状态  - ON：开启  - OFF：关闭
	Status string `json:"status"`
}

func (o SwitchStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchStatus struct{}"
	}

	return strings.Join([]string{"SwitchStatus", string(data)}, " ")
}
