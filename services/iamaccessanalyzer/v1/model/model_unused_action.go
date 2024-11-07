package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UnusedAction struct {

	// Action 名称。
	Action string `json:"action"`
}

func (o UnusedAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnusedAction struct{}"
	}

	return strings.Join([]string{"UnusedAction", string(data)}, " ")
}
