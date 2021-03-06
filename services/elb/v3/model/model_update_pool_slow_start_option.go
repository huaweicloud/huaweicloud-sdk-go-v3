package model

import (
	"encoding/json"

	"strings"
)

// 慢启动信息
type UpdatePoolSlowStartOption struct {
	// 慢启动的开关，默认值：false； true：开启； false：关闭

	Enable bool `json:"enable"`
	// 慢启动的持续时间，单位：s。默认：30； 取值范围：30~1200

	Duration int32 `json:"duration"`
}

func (o UpdatePoolSlowStartOption) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePoolSlowStartOption struct{}"
	}

	return strings.Join([]string{"UpdatePoolSlowStartOption", string(data)}, " ")
}
