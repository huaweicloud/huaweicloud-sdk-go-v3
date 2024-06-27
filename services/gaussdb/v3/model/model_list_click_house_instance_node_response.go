package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClickHouseInstanceNodeResponse Response Object
type ListClickHouseInstanceNodeResponse struct {

	// ClickHouse实例节点列表。
	NodeList       *[]ClickHouseNodeInfoResponseBodyNodeList `json:"node_list,omitempty"`
	HttpStatusCode int                                       `json:"-"`
}

func (o ListClickHouseInstanceNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClickHouseInstanceNodeResponse struct{}"
	}

	return strings.Join([]string{"ListClickHouseInstanceNodeResponse", string(data)}, " ")
}
