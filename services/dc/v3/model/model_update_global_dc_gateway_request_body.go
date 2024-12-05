package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGlobalDcGatewayRequestBody 创建全球接入网关请求体
type UpdateGlobalDcGatewayRequestBody struct {

	// 空运行？ true-是，false-否
	DryRun *bool `json:"dry_run,omitempty"`

	GlobalDcGateway *UpdateGlobalDcGatewayRequestBodyGlobalDcGateway `json:"global_dc_gateway,omitempty"`
}

func (o UpdateGlobalDcGatewayRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGlobalDcGatewayRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateGlobalDcGatewayRequestBody", string(data)}, " ")
}
