package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AntivirusEventType **参数解释**: 病毒查杀结果对应的事件类型标识 **取值范围**: 0-10（具体含义：0=文件病毒事件、1=内存病毒事件...，详见产品错误码/枚举文档）
type AntivirusEventType struct {
}

func (o AntivirusEventType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AntivirusEventType struct{}"
	}

	return strings.Join([]string{"AntivirusEventType", string(data)}, " ")
}
