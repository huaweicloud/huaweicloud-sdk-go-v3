package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCentralNetworkCapabilitiesRequest Request Object
type ListCentralNetworkCapabilitiesRequest struct {

	// 根据租户能力名查询，可查询多个类型。
	Capability *[]CentralNetworkCapabilityEnum `json:"capability,omitempty"`
}

func (o ListCentralNetworkCapabilitiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCentralNetworkCapabilitiesRequest struct{}"
	}

	return strings.Join([]string{"ListCentralNetworkCapabilitiesRequest", string(data)}, " ")
}
