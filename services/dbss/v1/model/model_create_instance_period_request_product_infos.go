package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateInstancePeriodRequestProductInfos struct {

	// 产品ID
	ProductId string `json:"product_id"`

	// 服务类型： 默认hws.service.type.dbss
	CloudServiceType string `json:"cloud_service_type"`

	// 资源类型: 默认hws.resource.type.dbss
	ResourceType string `json:"resource_type"`

	// 资源规格： dbss.bypassaudit.low、dbss.bypassaudit.medium、dbss.bypassaudit.high
	ResourceSpecCode string `json:"resource_spec_code"`

	// 产品规格描述
	ProductSpecDesc string `json:"product_spec_desc"`
}

func (o CreateInstancePeriodRequestProductInfos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstancePeriodRequestProductInfos struct{}"
	}

	return strings.Join([]string{"CreateInstancePeriodRequestProductInfos", string(data)}, " ")
}
