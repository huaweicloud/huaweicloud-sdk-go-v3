package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAvailabilityZoneInfosResponse Response Object
type ListAvailabilityZoneInfosResponse struct {

	// 可用区集合
	Azs            *[]AzInfo `json:"azs,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAvailabilityZoneInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailabilityZoneInfosResponse struct{}"
	}

	return strings.Join([]string{"ListAvailabilityZoneInfosResponse", string(data)}, " ")
}
