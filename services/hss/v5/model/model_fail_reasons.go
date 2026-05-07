package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FailReasons 主机结果
type FailReasons struct {

	// **参数解释** 原因 **取值范围** 长度1-100
	Reason *string `json:"reason,omitempty"`

	// **参数解释** 主机ID **取值范围** 长度1-64
	HostId *string `json:"host_id,omitempty"`
}

func (o FailReasons) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FailReasons struct{}"
	}

	return strings.Join([]string{"FailReasons", string(data)}, " ")
}
