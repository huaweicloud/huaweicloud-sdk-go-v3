package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeKillTaskSwitchRequestBody 设置自治限流开关请求体
type ChangeKillTaskSwitchRequestBody struct {

	// 开关状态
	Status string `json:"status"`
}

func (o ChangeKillTaskSwitchRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeKillTaskSwitchRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeKillTaskSwitchRequestBody", string(data)}, " ")
}
