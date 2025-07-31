package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ContainerId **参数解释**: 容器ID **取值范围**: 字符长度1-128位
type ContainerId struct {
}

func (o ContainerId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerId struct{}"
	}

	return strings.Join([]string{"ContainerId", string(data)}, " ")
}
