package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GlobalDcGatewayProjectId Gdgw的项目ID。
type GlobalDcGatewayProjectId struct {

	// Gdgw的项目ID。
	GlobalDcGatewayProjectId string `json:"global_dc_gateway_project_id"`
}

func (o GlobalDcGatewayProjectId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlobalDcGatewayProjectId struct{}"
	}

	return strings.Join([]string{"GlobalDcGatewayProjectId", string(data)}, " ")
}
