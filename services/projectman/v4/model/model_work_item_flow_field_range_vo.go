package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkItemFlowFieldRangeVo 工作项字段范围
type WorkItemFlowFieldRangeVo struct {

	// 可选值对象列表
	SettingValObject *[]map[string]interface{} `json:"setting_val_object,omitempty"`
}

func (o WorkItemFlowFieldRangeVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkItemFlowFieldRangeVo struct{}"
	}

	return strings.Join([]string{"WorkItemFlowFieldRangeVo", string(data)}, " ")
}
