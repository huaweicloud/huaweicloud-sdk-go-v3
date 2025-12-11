package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OperateResultRequestInfo 处置的结果
type OperateResultRequestInfo struct {

	// **参数解释**: 主机上安装的杀毒Agent的唯一标识ID，用于关联主机与杀毒服务 **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	AgentId string `json:"agent_id"`

	// **参数解释**： 病毒查杀结果ID **取值范围**： 字符长度1-64位
	ResultId string `json:"result_id"`

	// **参数解释**: 病毒查杀结果对应的事件类型标识 **取值范围**: 0-10（具体含义：0=文件病毒事件、1=内存病毒事件...，详见产品错误码/枚举文档）
	EventType int32 `json:"event_type"`

	// **参数解释**： 发生时间，毫秒 **取值范围**： 最小值0，最大值9223372036854775807，时间格式：毫秒级时间戳（UTC时区，从1970-01-01 00:00:00开始计算），单位：ms
	OccurTime *int64 `json:"occur_time,omitempty"`

	// **参数解释**： 文件哈希 **取值范围**： 字符长度1-256位
	FileHash string `json:"file_hash"`

	// **参数解释**： 文件路径 **取值范围**： 字符长度1-256位
	FilePath string `json:"file_path"`

	// **参数解释**： 文件的系统属性（如读写权限、隐藏属性、执行权限等） **取值范围**： 字符长度1-256位
	FileAttr string `json:"file_attr"`
}

func (o OperateResultRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateResultRequestInfo struct{}"
	}

	return strings.Join([]string{"OperateResultRequestInfo", string(data)}, " ")
}
