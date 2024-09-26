package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GlobalDcGatewayRegionId Gdgw的RegionID。
type GlobalDcGatewayRegionId struct {

	// Gdgw的RegionID。
	GlobalDcGatewayRegionId string `json:"global_dc_gateway_region_id"`
}

func (o GlobalDcGatewayRegionId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlobalDcGatewayRegionId struct{}"
	}

	return strings.Join([]string{"GlobalDcGatewayRegionId", string(data)}, " ")
}
