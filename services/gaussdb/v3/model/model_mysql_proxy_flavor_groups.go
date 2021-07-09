package model

import (
	"encoding/json"

	"strings"
)

type MysqlProxyFlavorGroups struct {
	// 规格组类型,如x86，arm。

	GroupType *string `json:"group_type,omitempty"`
	// 规格信息。

	ProxyFlavors *[]MysqlProxyComputeFlavor `json:"proxy_flavors,omitempty"`
}

func (o MysqlProxyFlavorGroups) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MysqlProxyFlavorGroups struct{}"
	}

	return strings.Join([]string{"MysqlProxyFlavorGroups", string(data)}, " ")
}
