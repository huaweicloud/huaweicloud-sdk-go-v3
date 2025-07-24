package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerHardwareAttributesRequest Request Object
type ShowServerHardwareAttributesRequest struct {

	// imetal server id
	Id string `json:"id"`
}

func (o ShowServerHardwareAttributesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerHardwareAttributesRequest struct{}"
	}

	return strings.Join([]string{"ShowServerHardwareAttributesRequest", string(data)}, " ")
}
