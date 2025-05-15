package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudConnectionCapabilitiesResponse Response Object
type ListCloudConnectionCapabilitiesResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	// 租户能力列表。
	Capabilities   []CloudConnectionCapability `json:"capabilities"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListCloudConnectionCapabilitiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudConnectionCapabilitiesResponse struct{}"
	}

	return strings.Join([]string{"ListCloudConnectionCapabilitiesResponse", string(data)}, " ")
}
