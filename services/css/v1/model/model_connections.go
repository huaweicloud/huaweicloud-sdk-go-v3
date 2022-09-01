package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 连接信息。
type Connections struct {

	// 终端节点ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 终端节点状态。 - accepted：允许该终端节点连接。 - rejected：拒绝该终端节点连接。
	Status *string `json:"status,omitempty" xml:"status"`

	// 最大连接数。
	MaxSession *string `json:"maxSession,omitempty" xml:"maxSession"`

	// 终端节点名称。
	SpecificationName *string `json:"specificationName,omitempty" xml:"specificationName"`

	// 创建时间，格式为ISO8601：CCYY-MM-DDThh:mm:ss。
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at"`

	// 更新时间。默认为null。
	UpdateAt *string `json:"update_at,omitempty" xml:"update_at"`

	// 拥有者。
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id"`
}

func (o Connections) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Connections struct{}"
	}

	return strings.Join([]string{"Connections", string(data)}, " ")
}
