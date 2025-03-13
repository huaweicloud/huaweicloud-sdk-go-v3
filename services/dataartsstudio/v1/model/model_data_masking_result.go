package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataMaskingResult 数据脱敏诊断结果
type DataMaskingResult struct {
	Result *DiagnoseResult `json:"result,omitempty"`

	// 没有配置脱敏任务的表数量
	Count *int32 `json:"count,omitempty"`
}

func (o DataMaskingResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataMaskingResult struct{}"
	}

	return strings.Join([]string{"DataMaskingResult", string(data)}, " ")
}
