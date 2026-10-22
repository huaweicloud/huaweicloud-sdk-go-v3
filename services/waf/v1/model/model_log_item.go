package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogItem 检测日志项
type LogItem struct {

	// 日志唯一标识
	Id *string `json:"id,omitempty"`

	// 检测时间（毫秒时间戳）
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 引擎ID
	EngineId *string `json:"engine_id,omitempty"`

	// APP ID
	AppId *string `json:"app_id,omitempty"`

	// APP名称
	AppName *string `json:"app_name,omitempty"`

	// 租户ID
	TenantId *string `json:"tenant_id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 检测类型
	DetectType *string `json:"detect_type,omitempty"`

	// 检测内容原文
	DetectContent *string `json:"detect_content,omitempty"`

	// 多轮对话历史上下文，与检测内容原文一一对应
	History *[]HistoryItem `json:"history,omitempty"`

	// 检测分类
	AttackTypes *[]string `json:"attack_types,omitempty"`

	// 执行动作
	Suggestion *string `json:"suggestion,omitempty"`

	// 安全代答
	CustomResp *string `json:"custom_resp,omitempty"`

	// 检测命中详情
	HitDetails *[]HitDetail `json:"hit_details,omitempty"`
}

func (o LogItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogItem struct{}"
	}

	return strings.Join([]string{"LogItem", string(data)}, " ")
}
