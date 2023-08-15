package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecretsRequest Request Object
type ListSecretsRequest struct {

	// 查询的起始位置，取值范围为非负整数，默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量，取值范围1~100，默认为100
	Limit *int32 `json:"limit,omitempty"`

	// 设备名称，模糊匹配，只允许中文字符、英文字母、数字、下划线、中划线，最大长度64
	Name *string `json:"name,omitempty"`

	// 工作空间ID，默认为注册账号子账号的default工作空间，可通过专业版HiLens控制台展开工作空间列表获取到工作空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 标签的key和value通过点连接，多个标签通过逗号连接，如：tags=key1.value1,key2.value2
	Tags *string `json:"tags,omitempty"`

	// 服务提供者：ief或hilens。不传会查询全部服务类型的设备列表
	Provider *string `json:"provider,omitempty"`

	// 排序方式，可根据名称、创建时间、更新时间排序枚举值：name，created_at，updated_at。sort默认升序，如sort=name，降序：sort=name%3Adesc。不填默认为sort=created_at%3Adesc。
	Sort *string `json:"sort,omitempty"`
}

func (o ListSecretsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecretsRequest struct{}"
	}

	return strings.Join([]string{"ListSecretsRequest", string(data)}, " ")
}
