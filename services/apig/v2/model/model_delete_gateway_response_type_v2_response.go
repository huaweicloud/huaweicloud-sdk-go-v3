package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGatewayResponseTypeV2Response Response Object
type DeleteGatewayResponseTypeV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteGatewayResponseTypeV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGatewayResponseTypeV2Response struct{}"
	}

	return strings.Join([]string{"DeleteGatewayResponseTypeV2Response", string(data)}, " ")
}
