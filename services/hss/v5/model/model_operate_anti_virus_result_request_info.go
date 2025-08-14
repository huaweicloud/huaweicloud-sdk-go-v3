package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OperateAntiVirusResultRequestInfo 处置病毒扫描结果
type OperateAntiVirusResultRequestInfo struct {

	// 处理方式，包含如下:   - mark_as_handled：手动处理   - ignore：忽略   - add_to_alarm_whitelist：加入告警白名单   - manual_isolate_and_kill：隔离文件   - unhandle：取消手动处理   - do_not_ignore：取消忽略   - remove_from_alarm_whitelist：删除告警白名单   - do_not_isolate_or_kill：取消隔离文件
	OperateType string `json:"operate_type"`

	// 备注信息
	Memo *string `json:"memo,omitempty"`

	// 处置的结果列表
	OperateResults *[]OperateResultRequestInfo `json:"operate_results,omitempty"`

	// 新增告警白名单规则列表
	EventWhiteRules *[]AntiVirusEventWhiteRuleListRequestInfo `json:"event_white_rules,omitempty"`
}

func (o OperateAntiVirusResultRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateAntiVirusResultRequestInfo struct{}"
	}

	return strings.Join([]string{"OperateAntiVirusResultRequestInfo", string(data)}, " ")
}
