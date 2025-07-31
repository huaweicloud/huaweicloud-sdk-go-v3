package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProcessEventName 事件名称   - \"white list alert malicious process\"   - \"white list alert suspicious process\"
type ProcessEventName struct {
}

func (o ProcessEventName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessEventName struct{}"
	}

	return strings.Join([]string{"ProcessEventName", string(data)}, " ")
}
