package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDataConnectorResponse Response Object
type CreateDataConnectorResponse struct {

	// 数据连接创建成功后系统返回的数据连接ID值。
	ConnectorId    *string `json:"connector_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDataConnectorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataConnectorResponse struct{}"
	}

	return strings.Join([]string{"CreateDataConnectorResponse", string(data)}, " ")
}
