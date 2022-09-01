package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ES扫描结果信息
type EsScanResultInfo struct {

	// 任务ID
	TaskId *string `json:"task_id,omitempty" xml:"task_id"`

	// 索引名
	IndexName *string `json:"index_name,omitempty" xml:"index_name"`

	// 类型ID
	TypeId *string `json:"type_id,omitempty" xml:"type_id"`

	// 类型名
	TypeName *string `json:"type_name,omitempty" xml:"type_name"`

	// 风险等级
	RiskLevel *int32 `json:"risk_level,omitempty" xml:"risk_level"`

	// 敏感数据类型
	SensitiveDataType *[]string `json:"sensitive_data_type,omitempty" xml:"sensitive_data_type"`

	// 规则匹配具体信息
	MatchInfo *[]EsMatchInfo `json:"match_info,omitempty" xml:"match_info"`
}

func (o EsScanResultInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EsScanResultInfo struct{}"
	}

	return strings.Join([]string{"EsScanResultInfo", string(data)}, " ")
}
