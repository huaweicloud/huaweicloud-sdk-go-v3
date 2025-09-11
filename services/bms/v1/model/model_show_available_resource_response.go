package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAvailableResourceResponse Response Object
type ShowAvailableResourceResponse struct {
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	Flavors        *[]FlavorResource `json:"flavors,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowAvailableResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAvailableResourceResponse struct{}"
	}

	return strings.Join([]string{"ShowAvailableResourceResponse", string(data)}, " ")
}
