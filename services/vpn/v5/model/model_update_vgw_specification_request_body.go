package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateVgwSpecificationRequestBody struct {
	VpnGateway *UpdateVgwSpecificationRequestBodyContent `json:"vpn_gateway"`
}

func (o UpdateVgwSpecificationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVgwSpecificationRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateVgwSpecificationRequestBody", string(data)}, " ")
}
