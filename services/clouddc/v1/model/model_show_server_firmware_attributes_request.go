package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerFirmwareAttributesRequest Request Object
type ShowServerFirmwareAttributesRequest struct {

	// imetal server id
	Id string `json:"id"`
}

func (o ShowServerFirmwareAttributesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerFirmwareAttributesRequest struct{}"
	}

	return strings.Join([]string{"ShowServerFirmwareAttributesRequest", string(data)}, " ")
}
