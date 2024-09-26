package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCentralNetworkCapabilitiesResponse Response Object
type ListCentralNetworkCapabilitiesResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	// 租户能力列表
	Capabilities   []CentralNetworkCapability `json:"capabilities"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListCentralNetworkCapabilitiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCentralNetworkCapabilitiesResponse struct{}"
	}

	return strings.Join([]string{"ListCentralNetworkCapabilitiesResponse", string(data)}, " ")
}
