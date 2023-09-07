package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ActionBaseInfo Connector执行动作基本数据
type ActionBaseInfo struct {

	// 执行动作的类型
	ActionType *interface{} `json:"action_type,omitempty"`

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

	// 执行动作ID
	Id *string `json:"id,omitempty"`

	// 执行动作名称
	Name *string `json:"name,omitempty"`

	// 执行动作ID
	OperationId *string `json:"operation_id,omitempty"`

	// swagger文档
	Swagger *interface{} `json:"swagger,omitempty"`

	// 最近一次测试结果
	TestResult *string `json:"test_result,omitempty"`

	// 修改时间
	UpdatedTime *sdktime.SdkTime `json:"updated_time,omitempty"`

	// 标记动作在流编排是否可见
	Visibility *string `json:"visibility,omitempty"`
}

func (o ActionBaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionBaseInfo struct{}"
	}

	return strings.Join([]string{"ActionBaseInfo", string(data)}, " ")
}
