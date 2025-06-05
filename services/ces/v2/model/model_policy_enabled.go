package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyEnabled 是否启用告警策略。true:开启，false:关闭。
type PolicyEnabled struct {
}

func (o PolicyEnabled) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyEnabled struct{}"
	}

	return strings.Join([]string{"PolicyEnabled", string(data)}, " ")
}
