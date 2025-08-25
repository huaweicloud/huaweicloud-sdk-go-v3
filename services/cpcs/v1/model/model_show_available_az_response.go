package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAvailableAzResponse Response Object
type ShowAvailableAzResponse struct {

	// 可用区列表
	AvailabilityZone *[]ShowAvailableAzResponsebodyAvailabilityZone `json:"availability_zone,omitempty"`
	HttpStatusCode   int                                            `json:"-"`
}

func (o ShowAvailableAzResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAvailableAzResponse struct{}"
	}

	return strings.Join([]string{"ShowAvailableAzResponse", string(data)}, " ")
}
