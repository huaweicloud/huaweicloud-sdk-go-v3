package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGlobalDcGatewayRequestBody 更新global-dc-gateway的请求体。
type UpdateGlobalDcGatewayRequestBody struct {
	GlobalDcGateway *UpdateGlobalDcGateway `json:"global_dc_gateway"`
}

func (o UpdateGlobalDcGatewayRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGlobalDcGatewayRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateGlobalDcGatewayRequestBody", string(data)}, " ")
}
