package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClassificationResult 数据分级分类诊断结果
type ClassificationResult struct {
	Result *DiagnoseResult `json:"result,omitempty"`

	// 是否配置了密级
	SecurityLevel *bool `json:"security_level,omitempty"`

	// 是否配置了分类
	Classification *bool `json:"classification,omitempty"`
}

func (o ClassificationResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClassificationResult struct{}"
	}

	return strings.Join([]string{"ClassificationResult", string(data)}, " ")
}
