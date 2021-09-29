package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowGaussMySqlProxyResponse struct {
	Proxy *MysqlProxy `json:"proxy,omitempty"`

	MasterNode *MysqlProxyNode `json:"master_node,omitempty"`
	// 只读节点信息。

	ReadonlyNodes  *[]MysqlProxyNode `json:"readonly_nodes,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowGaussMySqlProxyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowGaussMySqlProxyResponse struct{}"
	}

	return strings.Join([]string{"ShowGaussMySqlProxyResponse", string(data)}, " ")
}
