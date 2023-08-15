package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowInstanceExtendProductInfoRespHourly struct {

	// 消息引擎的名称，该字段显示为rabbitmq。
	Name *string `json:"name,omitempty"`

	// 消息引擎的版本，当前支持3.8.35[和3.7.17](tag:hk_sbc,sbc)。
	Version *string `json:"version,omitempty"`

	// 产品规格列表。
	Values *[]ListProductsRespValues `json:"values,omitempty"`
}

func (o ShowInstanceExtendProductInfoRespHourly) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceExtendProductInfoRespHourly struct{}"
	}

	return strings.Join([]string{"ShowInstanceExtendProductInfoRespHourly", string(data)}, " ")
}
