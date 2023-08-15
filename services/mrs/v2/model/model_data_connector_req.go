package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataConnectorReq 创建数据连接请求
type DataConnectorReq struct {
	DataConnector *DataConnector `json:"data_connector"`
}

func (o DataConnectorReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataConnectorReq struct{}"
	}

	return strings.Join([]string{"DataConnectorReq", string(data)}, " ")
}
