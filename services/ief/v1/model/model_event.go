package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Event 系统订阅事件基本信息
type Event struct {

	// 系统订阅事件ID
	Id *string `json:"id,omitempty"`

	// 系统订阅事件名称。只允许小写英文字符、数字、下划线、中划线，最大长度64，同一个账号中创建的系统订阅和消息规则名唯一
	Name *string `json:"name,omitempty"`

	// 系统订阅事件所属项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// 描述，最大长度255，不允许^~#$%&*<>()[]{}'\"\\
	Description *string `json:"description,omitempty"`

	// 是否启用系统订阅规则，默认为true（启用）
	InUsing *bool `json:"in_using,omitempty"`

	// 系统订阅事件主题。每个主题由“{边缘资源}/{操作}”组成，多个主题使用逗号（,）进行分隔，支持如下主题： - edgeNode/offline：节点离线 - edgeNode/online：节点上线 - edgeNode/all：节点离线+节点上线 - deployment/created：容器应用创建 - deployment/updated：容器应用更新 - deployment/deleted：容器应用删除 - deployment/all：容器应用创建+更新+删除 - instance/created：应用实例创建 - instance/updated：应用实例更新 - instance/deleted：应用实例删除 - instance/all：应用实例创建+更新+删除
	Events *string `json:"events,omitempty"`

	Target *EndpointObjResp `json:"target,omitempty"`

	// 目的端点资源属性
	TargetResource map[string]string `json:"target_resource,omitempty"`

	// 成功次数
	SuccessMessages *int32 `json:"success_messages,omitempty"`

	// 失败次数
	FailMessages *int32 `json:"fail_messages,omitempty"`

	// 删除时间
	DeleteAt *int32 `json:"delete_at,omitempty"`
}

func (o Event) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Event struct{}"
	}

	return strings.Join([]string{"Event", string(data)}, " ")
}
