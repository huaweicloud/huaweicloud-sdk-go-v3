package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNodesRequest Request Object
type ShowNodesRequest struct {

	// 查询的起始位置，取值范围为非负整数，默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量，取值范围1~100，默认为100
	Limit *int32 `json:"limit,omitempty"`

	// 设备名称，模糊匹配，只允许中文字符、英文字母、数字、下划线、中划线，最大长度64
	Name *string `json:"name,omitempty"`

	// 工作空间ID，默认为注册账号/子账号的default工作空间，可通过专业版HiLens控制台展开工作空间列表获取到工作空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 应用名称，模糊匹配，只允许英文小写字母、数字、中划线，最大长度32, 英文小写字母或数字开头和结尾
	AppName *string `json:"app_name,omitempty"`

	// 标签的key和value通过点连接，多个标签通过逗号连接，如：tags=key1.value1,key2.value2
	Tags *string `json:"tags,omitempty"`

	// 服务提供者：ief或hilens。不传会查询全部服务类型的设备列表
	Provider *string `json:"provider,omitempty"`

	// 集群ID，若值为0会过滤出不隶属任何集群的设备
	ClusterId *string `json:"cluster_id,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 设备激活状态，分别是未激活（INACTIVE）、已激活（ACTIVATED）、已到期（EXPIRED）
	ActiveStatus *string `json:"active_status,omitempty"`
}

func (o ShowNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNodesRequest struct{}"
	}

	return strings.Join([]string{"ShowNodesRequest", string(data)}, " ")
}
