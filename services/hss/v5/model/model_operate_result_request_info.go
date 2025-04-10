package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OperateResultRequestInfo 处置的结果
type OperateResultRequestInfo struct {

	// Agent ID
	AgentId string `json:"agent_id"`

	// 病毒查杀结果ID
	ResultId string `json:"result_id"`

	// 事件类型
	EventType int32 `json:"event_type"`

	// 发生时间，毫秒
	OccurTime *int64 `json:"occur_time,omitempty"`

	// 文件哈希
	FileHash string `json:"file_hash"`

	// 文件路径
	FilePath string `json:"file_path"`

	// 文件属性
	FileAttr string `json:"file_attr"`
}

func (o OperateResultRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateResultRequestInfo struct{}"
	}

	return strings.Join([]string{"OperateResultRequestInfo", string(data)}, " ")
}
