package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCustomConnectorsResponse Response Object
type ShowCustomConnectorsResponse struct {

	// 连接器列表
	ConnectorInfos *[]ConnectorInfo0 `json:"connector_infos,omitempty"`

	// 连接器数量
	Count          *int64 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowCustomConnectorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCustomConnectorsResponse struct{}"
	}

	return strings.Join([]string{"ShowCustomConnectorsResponse", string(data)}, " ")
}
