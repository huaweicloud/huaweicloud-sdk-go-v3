package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataConnectorResponse Response Object
type ListDataConnectorResponse struct {

	// 数据连接总数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 数据连接详情列表
	DataConnectors *[]DataConnectorDetail `json:"data_connectors,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListDataConnectorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataConnectorResponse struct{}"
	}

	return strings.Join([]string{"ListDataConnectorResponse", string(data)}, " ")
}
