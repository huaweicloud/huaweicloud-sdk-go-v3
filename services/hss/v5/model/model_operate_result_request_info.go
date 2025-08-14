package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OperateResultRequestInfo 处置的结果
type OperateResultRequestInfo struct {

	// **参数解释**: Agent ID **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	AgentId string `json:"agent_id"`

	// 病毒查杀结果ID
	ResultId string `json:"result_id"`

	// 事件类型
	EventType int32 `json:"event_type"`

	// **参数解释**： 发生时间，毫秒 **取值范围**： 最小值0，最大值9223372036854775807
	OccurTime *int64 `json:"occur_time,omitempty"`

	// **参数解释**： 文件哈希 **取值范围**： 字符长度1-256位
	FileHash string `json:"file_hash"`

	// **参数解释**： 文件路径 **取值范围**： 字符长度1-256位
	FilePath string `json:"file_path"`

	// **参数解释**： 文件属性 **取值范围**： 字符长度1-256位
	FileAttr string `json:"file_attr"`
}

func (o OperateResultRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateResultRequestInfo struct{}"
	}

	return strings.Join([]string{"OperateResultRequestInfo", string(data)}, " ")
}
