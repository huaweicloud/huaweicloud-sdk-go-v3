package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HistoryHandleMethod 处理方式，包含如下:   - mark_as_handled : 手动处理   - ignore : 忽略   - add_to_alarm_whitelist : 加入告警白名单   - isolate_and_kill : 隔离文件
type HistoryHandleMethod struct {
}

func (o HistoryHandleMethod) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HistoryHandleMethod struct{}"
	}

	return strings.Join([]string{"HistoryHandleMethod", string(data)}, " ")
}
