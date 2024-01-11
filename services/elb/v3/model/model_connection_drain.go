package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConnectionDrain struct {

	// 延迟注销功能开关，默认值：false； true：开启； false：关闭
	Enable *bool `json:"enable,omitempty"`

	// 延迟注销时间，单位：s； 取值范围：10~4000
	Timeout *int32 `json:"timeout,omitempty"`
}

func (o ConnectionDrain) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionDrain struct{}"
	}

	return strings.Join([]string{"ConnectionDrain", string(data)}, " ")
}
