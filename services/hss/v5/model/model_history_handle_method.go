package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HistoryHandleMethod **参数解释**: 处理方式 **取值范围**: 包含如下:   - mark_as_handled：手动处理   - ignore：忽略   - add_to_alarm_whitelist：加入告警白名单   - isolate_and_kill：隔离文件   - unhandle：取消手动处理   - do_not_ignore：取消忽略   - remove_from_alarm_whitelist：删除告警白名单   - do_not_isolate_or_kill：取消隔离文件
type HistoryHandleMethod struct {
}

func (o HistoryHandleMethod) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HistoryHandleMethod struct{}"
	}

	return strings.Join([]string{"HistoryHandleMethod", string(data)}, " ")
}
