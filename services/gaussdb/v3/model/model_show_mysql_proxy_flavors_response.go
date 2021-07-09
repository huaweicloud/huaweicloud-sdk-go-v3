package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowMysqlProxyFlavorsResponse struct {
	// 规格组信息。

	ProxyFlavorGroups *[]MysqlProxyFlavorGroups `json:"proxy_flavor_groups,omitempty"`
	HttpStatusCode    int                       `json:"-"`
}

func (o ShowMysqlProxyFlavorsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowMysqlProxyFlavorsResponse struct{}"
	}

	return strings.Join([]string{"ShowMysqlProxyFlavorsResponse", string(data)}, " ")
}
