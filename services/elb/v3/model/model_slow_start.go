package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SlowStart struct {

	// **参数解释**：慢启动的开关。  **取值范围**： - true：开启。 - false：关闭。
	Enable bool `json:"enable"`

	// **参数解释**：慢启动的持续时间。  **取值范围**：30-1200，单位：秒。
	Duration int32 `json:"duration"`
}

func (o SlowStart) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowStart struct{}"
	}

	return strings.Join([]string{"SlowStart", string(data)}, " ")
}
