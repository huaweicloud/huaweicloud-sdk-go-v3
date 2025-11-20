package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageSensitiveInfo struct {

	// **参数解释**: 敏感事件编号 **取值范围**: 字符长度0-128位
	SensitiveInfoId *string `json:"sensitive_info_id,omitempty"`

	// **参数解释**: 威胁等级 **取值范围**: - critical：致命。 - high：高危。 - medium：中危。 - low : 低危。
	Severity *string `json:"severity,omitempty"`

	// **参数解释**: 规则名称 **取值范围**: 字符长度0-128位
	Name *string `json:"name,omitempty"`

	// **参数解释**: 规则描述 **取值范围**: 字符长度0-128位
	Description *string `json:"description,omitempty"`

	// **参数解释**: 漏洞所在镜像层 **取值范围**: 字符长度0-128位
	Position *string `json:"position,omitempty"`

	// **参数解释**: 文件路径 **取值范围**: 字符长度0-128位
	FilePath *string `json:"file_path,omitempty"`

	// **参数解释**: 敏感信息内容 **取值范围**: 字符长度0-128位
	Content *string `json:"content,omitempty"`

	// **参数解释**: 最后一次检测时间，时间单位 毫秒（ms） **取值范围**: 最小值0，最大值2147483647
	LatestScanTime *int64 `json:"latest_scan_time,omitempty"`

	// **参数解释**: 是否已处理 **取值范围**: - unhandled：未处理。 - handled：已处理。
	HandleStatus *string `json:"handle_status,omitempty"`

	// **参数解释**: 操作类型 **取值范围**: - ignore ：忽略。 - do_not_ignore ：取消忽略。
	OperateAccept *string `json:"operate_accept,omitempty"`
}

func (o ImageSensitiveInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageSensitiveInfo struct{}"
	}

	return strings.Join([]string{"ImageSensitiveInfo", string(data)}, " ")
}
