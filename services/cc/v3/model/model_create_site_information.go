package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSiteInformation struct {

	// RegionID。
	RegionId string `json:"region_id"`

	// 实例所属项目ID。
	ProjectId string `json:"project_id"`

	GatewayType *GatewayTypeEnum `json:"gateway_type"`

	// 网关的ID。
	GatewayId string `json:"gateway_id"`
}

func (o CreateSiteInformation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSiteInformation struct{}"
	}

	return strings.Join([]string{"CreateSiteInformation", string(data)}, " ")
}
