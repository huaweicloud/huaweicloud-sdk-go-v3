package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 连接信息。
type Connections struct {
	// 终端节点ID。

	Id *string `json:"id,omitempty"`
	// 状态。

	Status *string `json:"status,omitempty"`
	// 最大连接数。

	MaxSession *string `json:"maxSession,omitempty"`
	// 终端节点名称。

	SpecificationName *string `json:"specificationName,omitempty"`
	// 创建时间。

	CreatedAt *string `json:"created_at,omitempty"`
	// 更新时间。

	UpdateAt *string `json:"update_at,omitempty"`
	// 拥有者。

	DomainId *string `json:"domain_id,omitempty"`
}

func (o Connections) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Connections struct{}"
	}

	return strings.Join([]string{"Connections", string(data)}, " ")
}
