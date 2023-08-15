package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListProductsRespHourly struct {

	// 消息引擎的名称，该字段显示为rabbitmq。
	Name *string `json:"name,omitempty"`

	// 消息引擎的版本，当前支持3.8.35[和3.7.17](tag:sbc,hk_sbc)。
	Version *string `json:"version,omitempty"`

	// 产品规格列表。
	Values *[]ListProductsRespValues `json:"values,omitempty"`
}

func (o ListProductsRespHourly) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProductsRespHourly struct{}"
	}

	return strings.Join([]string{"ListProductsRespHourly", string(data)}, " ")
}
