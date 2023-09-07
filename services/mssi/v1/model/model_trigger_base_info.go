package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TriggerBaseInfo Connector触发事件基本数据
type TriggerBaseInfo struct {

	// 分类
	Category *string `json:"category,omitempty"`

	ConnectorActionHtml *string `json:"connector_action_html,omitempty"`

	ConnectorCreatedType *string `json:"connector_created_type,omitempty"`

	// 自定义连接器ID
	ConnectorId *string `json:"connector_id,omitempty"`

	// 连接器版本
	ConnectorVersion *string `json:"connector_version,omitempty"`

	// 创建时间
	CreatedTime *sdktime.SdkTime `json:"created_time,omitempty"`

	// 操作or触发器的详细定义
	Definition *interface{} `json:"definition,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 触发事件ID
	Id *string `json:"id,omitempty"`

	// 执行动作名称
	Name *string `json:"name,omitempty"`

	// 执行动作ID
	OperationId *string `json:"operation_id,omitempty"`

	// swagger文档
	Swagger *interface{} `json:"swagger,omitempty"`

	// 最近一次测试结果
	TestResult *string `json:"test_result,omitempty"`

	// 触发事件的类型
	TriggerType *interface{} `json:"trigger_type,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 修改时间
	UpdatedTime *sdktime.SdkTime `json:"updated_time,omitempty"`
}

func (o TriggerBaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TriggerBaseInfo struct{}"
	}

	return strings.Join([]string{"TriggerBaseInfo", string(data)}, " ")
}
