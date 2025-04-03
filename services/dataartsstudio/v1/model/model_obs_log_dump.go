package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ObsLogDump struct {

	// 是否开启Obs日志转储功能，true表示开启，false表示关闭。
	LogDump *bool `json:"log_dump,omitempty"`
}

func (o ObsLogDump) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsLogDump struct{}"
	}

	return strings.Join([]string{"ObsLogDump", string(data)}, " ")
}
