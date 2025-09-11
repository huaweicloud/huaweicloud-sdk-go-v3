package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAvailabilityZoneInfosNewResponse Response Object
type ListAvailabilityZoneInfosNewResponse struct {

	// 可用区集合
	Azs            *[]AzInfo `json:"azs,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAvailabilityZoneInfosNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailabilityZoneInfosNewResponse struct{}"
	}

	return strings.Join([]string{"ListAvailabilityZoneInfosNewResponse", string(data)}, " ")
}
