package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAvailableFlavorInfosResponse Response Object
type ListAvailableFlavorInfosResponse struct {

	// 实例id。
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称。
	InstanceName *string `json:"instance_name,omitempty"`

	CurrentFlavor *ComputeFlavor `json:"current_flavor,omitempty"`

	OptionalFlavors *OptionalFlavorsInfo `json:"optional_flavors,omitempty"`
	HttpStatusCode  int                  `json:"-"`
}

func (o ListAvailableFlavorInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableFlavorInfosResponse struct{}"
	}

	return strings.Join([]string{"ListAvailableFlavorInfosResponse", string(data)}, " ")
}
