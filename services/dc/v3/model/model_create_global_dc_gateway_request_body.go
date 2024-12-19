package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGlobalDcGatewayRequestBody 创建global-dc-gateway的请求体。
type CreateGlobalDcGatewayRequestBody struct {
	GlobalDcGateway *CreateGlobalDcGateway `json:"global_dc_gateway"`
}

func (o CreateGlobalDcGatewayRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGlobalDcGatewayRequestBody struct{}"
	}

	return strings.Join([]string{"CreateGlobalDcGatewayRequestBody", string(data)}, " ")
}
