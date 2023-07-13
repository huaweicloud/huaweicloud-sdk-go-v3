package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventCreateDetail 系统订阅创建配置
type EventCreateDetail struct {

	// 系统订阅名称。只允许小写英文字符、数字、下划线、中划线，最大长度64，同一个帐号中创建的系统订阅和消息规则名唯一
	Name *string `json:"name,omitempty"`

	// 描述。最大长度255，不允许^~#$%&*<>()[]{}'\"\\
	Description *string `json:"description,omitempty"`

	// 系统订阅主题。每个主题由“{边缘资源}/{操作}”组成，多个主题使用逗号（,）进行分隔，支持如下主题： - edgeNode/offline：节点离线 - edgeNode/online：节点上线 - edgeNode/all：节点离线+节点上线 - deployment/created：容器应用创建 - deployment/updated：容器应用更新 - deployment/deleted：容器应用删除 - deployment/all：容器应用创建+更新+删除 - instance/created：应用实例创建 - instance/updated：应用实例更新 - instance/deleted：应用实例删除 - instance/all：应用实例创建+更新+删除
	Events *string `json:"events,omitempty"`

	// 目的端点ID
	Target *string `json:"target,omitempty"`

	TargetResource *EndpointResource `json:"target_resource,omitempty"`
}

func (o EventCreateDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventCreateDetail struct{}"
	}

	return strings.Join([]string{"EventCreateDetail", string(data)}, " ")
}
