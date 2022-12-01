package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type License struct {
	ChannelLimit *int32 `json:"channel_limit,omitempty"`

	ChargingModel *int32 `json:"charging_model,omitempty"`

	CloudServiceType *string `json:"cloud_service_type,omitempty"`

	HilensResourceSpecCode *string `json:"hilens_resource_spec_code,omitempty"`

	MeasureType *string `json:"measure_type,omitempty"`

	MeasureUnit *string `json:"measure_unit,omitempty"`

	OfflineFlag *int32 `json:"offline_flag,omitempty"`

	OrderLimit *int32 `json:"order_limit,omitempty"`

	ProductInfo *[]ProductInfo `json:"product_info,omitempty"`

	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	ResourceStepSize *int32 `json:"resource_step_size,omitempty"`

	ResourceType *string `json:"resource_type,omitempty"`

	Validity *int32 `json:"validity,omitempty"`
}

func (o License) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "License struct{}"
	}

	return strings.Join([]string{"License", string(data)}, " ")
}
