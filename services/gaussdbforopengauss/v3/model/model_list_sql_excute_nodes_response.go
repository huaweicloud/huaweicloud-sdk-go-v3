package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSqlExcuteNodesResponse Response Object
type ListSqlExcuteNodesResponse struct {

	// **参数解释**: 节点信息列表。
	Nodes          *[]ExcuteSqlNodeInfoResult `json:"nodes,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListSqlExcuteNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlExcuteNodesResponse struct{}"
	}

	return strings.Join([]string{"ListSqlExcuteNodesResponse", string(data)}, " ")
}
