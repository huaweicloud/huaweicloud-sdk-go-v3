package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGlobalDcGatewayRequestBody 全球接入网关响应体
type CreateGlobalDcGatewayRequestBody struct {

	// 空运行 - true 是 - false 否
	DryRun *bool `json:"dry_run,omitempty"`

	GlobalDcGateway *CreateGlobalDcGatewayRequestBodyGlobalDcGateway `json:"global_dc_gateway,omitempty"`
}

func (o CreateGlobalDcGatewayRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGlobalDcGatewayRequestBody struct{}"
	}

	return strings.Join([]string{"CreateGlobalDcGatewayRequestBody", string(data)}, " ")
}
