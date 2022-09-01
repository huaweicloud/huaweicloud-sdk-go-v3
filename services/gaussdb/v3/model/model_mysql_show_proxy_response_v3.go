package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MysqlShowProxyResponseV3 struct {
	Proxy *MysqlProxyV3 `json:"proxy,omitempty" xml:"proxy"`

	MasterNode *MysqlProxyNodeV3 `json:"master_node,omitempty" xml:"master_node"`

	// 只读节点信息。
	ReadonlyNodes *[]MysqlProxyNodeV3 `json:"readonly_nodes,omitempty" xml:"readonly_nodes"`
}

func (o MysqlShowProxyResponseV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlShowProxyResponseV3 struct{}"
	}

	return strings.Join([]string{"MysqlShowProxyResponseV3", string(data)}, " ")
}
