package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MysqlProxyNodes struct {

	// Proxy节点id。
	Id *string `json:"id,omitempty" xml:"id"`

	// Proxy节点状态。 取值范围：ACTIVE、ABNORMAL、BUILD和FAILED。
	Status *string `json:"status,omitempty" xml:"status"`

	// Proxy节点名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// Proxy节点角色：master和slave。
	Role *string `json:"role,omitempty" xml:"role"`

	// 可用区。
	AzCode *string `json:"az_code,omitempty" xml:"az_code"`

	// Proxy节点是否被冻结：0-未冻结；1-冻结；2-冻结删除。
	FrozenFlag *int32 `json:"frozen_flag,omitempty" xml:"frozen_flag"`
}

func (o MysqlProxyNodes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlProxyNodes struct{}"
	}

	return strings.Join([]string{"MysqlProxyNodes", string(data)}, " ")
}
