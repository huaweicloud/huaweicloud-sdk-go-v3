package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDasCloudDbaPriceRequest Request Object
type ShowDasCloudDbaPriceRequest struct {
}

func (o ShowDasCloudDbaPriceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDasCloudDbaPriceRequest struct{}"
	}

	return strings.Join([]string{"ShowDasCloudDbaPriceRequest", string(data)}, " ")
}
