package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HookCreate struct {
	// 无法猜测的随机字符串，用于验证接收到的payloads。

	Secret string `json:"secret"`
	// hook触发时的回调URL。

	Url string `json:"url"`
}

func (o HookCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HookCreate struct{}"
	}

	return strings.Join([]string{"HookCreate", string(data)}, " ")
}
