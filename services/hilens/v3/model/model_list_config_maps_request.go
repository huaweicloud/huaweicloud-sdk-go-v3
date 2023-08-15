package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigMapsRequest Request Object
type ListConfigMapsRequest struct {

	// 服务提供者：ief或hilens，默认为hilens。
	Provider *string `json:"provider,omitempty"`

	// 配置项名称，模糊匹配。
	Name *string `json:"name,omitempty"`

	// 工作空间ID，默认为注册账号/子账号的default工作空间，可通过专业版HiLens控制台展开工作空间列表获取到工作空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 每页显示的条目数量，取值范围1~1000，默认值为1000。
	Limit *int32 `json:"limit,omitempty"`

	// 查询的起始位置，取值范围为非负整数，默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 排序方式，可根据名称、创建时间、更新时间排序枚举值：name，created_at，updated_at。sort默认升序，如sort=name，降序：sort=name%3Adesc。不填默认为sort=created_at%3Adesc。
	Sort *string `json:"sort,omitempty"`

	// 格式为{tag_key}={tag_value}，支持多对tag或查询。
	TagKey *string `json:"tag_key,omitempty"`
}

func (o ListConfigMapsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigMapsRequest struct{}"
	}

	return strings.Join([]string{"ListConfigMapsRequest", string(data)}, " ")
}
