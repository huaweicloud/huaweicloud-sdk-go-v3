package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConnectorsResponse Response Object
type ShowConnectorsResponse struct {

	// 连接器列表
	Connectors *[]ConnectorInfo `json:"connectors,omitempty"`

	// 连接器数量
	Count          *int64 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowConnectorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConnectorsResponse struct{}"
	}

	return strings.Join([]string{"ShowConnectorsResponse", string(data)}, " ")
}
