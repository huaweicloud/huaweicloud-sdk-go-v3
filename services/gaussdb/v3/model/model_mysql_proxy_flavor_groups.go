package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MysqlProxyFlavorGroups struct {

	// 规格组类型，如x86、arm。
	GroupType *string `json:"group_type,omitempty" xml:"group_type"`

	// 规格信息。
	ProxyFlavors *[]MysqlProxyComputeFlavor `json:"proxy_flavors,omitempty" xml:"proxy_flavors"`
}

func (o MysqlProxyFlavorGroups) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlProxyFlavorGroups struct{}"
	}

	return strings.Join([]string{"MysqlProxyFlavorGroups", string(data)}, " ")
}
