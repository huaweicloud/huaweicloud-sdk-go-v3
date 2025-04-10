package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AntivirusEventType 事件类型
type AntivirusEventType struct {
}

func (o AntivirusEventType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AntivirusEventType struct{}"
	}

	return strings.Join([]string{"AntivirusEventType", string(data)}, " ")
}
