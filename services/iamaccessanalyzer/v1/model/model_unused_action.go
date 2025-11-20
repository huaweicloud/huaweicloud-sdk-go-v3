package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UnusedAction struct {

	// 授权项名称。
	Action string `json:"action"`

	// 用户使用授权项的最后访问时间。
	LastAccessed *interface{} `json:"last_accessed,omitempty"`
}

func (o UnusedAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnusedAction struct{}"
	}

	return strings.Join([]string{"UnusedAction", string(data)}, " ")
}
