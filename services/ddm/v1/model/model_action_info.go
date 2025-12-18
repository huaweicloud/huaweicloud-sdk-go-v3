package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ActionInfo struct {

	// 锁。
	Action *string `json:"action,omitempty"`
}

func (o ActionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionInfo struct{}"
	}

	return strings.Join([]string{"ActionInfo", string(data)}, " ")
}
