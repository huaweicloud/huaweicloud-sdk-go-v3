package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventNum 事件数量
type EventNum struct {
}

func (o EventNum) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventNum struct{}"
	}

	return strings.Join([]string{"EventNum", string(data)}, " ")
}
