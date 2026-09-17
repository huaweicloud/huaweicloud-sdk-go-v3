package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkItemFlowFieldValueVo 工作项字段值
type WorkItemFlowFieldValueVo struct {

	// 引用属性名
	RefProp *string `json:"ref_prop,omitempty"`

	// 配置值对象列表
	SettingValObject *[]map[string]interface{} `json:"setting_val_object,omitempty"`
}

func (o WorkItemFlowFieldValueVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkItemFlowFieldValueVo struct{}"
	}

	return strings.Join([]string{"WorkItemFlowFieldValueVo", string(data)}, " ")
}
