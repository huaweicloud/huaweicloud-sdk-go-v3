package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EsScanResultInfo ES扫描结果信息
type EsScanResultInfo struct {

	// 任务ID
	TaskId *string `json:"task_id,omitempty"`

	// 索引名
	IndexName *string `json:"index_name,omitempty"`

	// 类型ID
	TypeId *string `json:"type_id,omitempty"`

	// 类型名
	TypeName *string `json:"type_name,omitempty"`

	// 风险等级
	RiskLevel *int32 `json:"risk_level,omitempty"`

	// 敏感数据类型
	SensitiveDataType *[]string `json:"sensitive_data_type,omitempty"`

	// 规则匹配具体信息
	MatchInfo *[]EsMatchInfo `json:"match_info,omitempty"`
}

func (o EsScanResultInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EsScanResultInfo struct{}"
	}

	return strings.Join([]string{"EsScanResultInfo", string(data)}, " ")
}
