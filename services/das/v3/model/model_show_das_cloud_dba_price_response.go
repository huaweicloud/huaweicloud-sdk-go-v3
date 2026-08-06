package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDasCloudDbaPriceResponse Response Object
type ShowDasCloudDbaPriceResponse struct {

	// 基础费用
	BasePrice *float64 `json:"base_price,omitempty"`

	// 基础费用-原价
	OriginalBasePrice *float64 `json:"original_base_price,omitempty"`

	// 存储费用
	StoragePrice *float64 `json:"storage_price,omitempty"`

	// 存储费用-原价
	OriginalStoragePrice *float64 `json:"original_storage_price,omitempty"`

	// 转储费用
	DumpPrice *float64 `json:"dump_price,omitempty"`

	// 转储费用-原价
	OriginalDumpPrice *float64 `json:"original_dump_price,omitempty"`

	// 度量单位标识,1:元
	MeasureId *int32 `json:"measure_id,omitempty"`

	// 币种，比如CNY
	Currency *string `json:"currency,omitempty"`

	// 配置费用-当前为0
	ConfigurePrice *float64 `json:"configure_price,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ShowDasCloudDbaPriceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDasCloudDbaPriceResponse struct{}"
	}

	return strings.Join([]string{"ShowDasCloudDbaPriceResponse", string(data)}, " ")
}
