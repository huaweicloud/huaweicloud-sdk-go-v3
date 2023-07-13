package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPublicIpTypeRequest Request Object
type ShowPublicIpTypeRequest struct {
}

func (o ShowPublicIpTypeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublicIpTypeRequest struct{}"
	}

	return strings.Join([]string{"ShowPublicIpTypeRequest", string(data)}, " ")
}
