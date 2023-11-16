package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFuncReservedInstanceMetricsResponse Response Object
type ShowFuncReservedInstanceMetricsResponse struct {

	// 弹性实例指标
	InstanceNum *[]SlaReportsValue `json:"instanceNum,omitempty"`

	// 预留实例指标
	ReservedInstanceNum *[]SlaReportsValue `json:"reservedInstanceNum,omitempty"`
	HttpStatusCode      int                `json:"-"`
}

func (o ShowFuncReservedInstanceMetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFuncReservedInstanceMetricsResponse struct{}"
	}

	return strings.Join([]string{"ShowFuncReservedInstanceMetricsResponse", string(data)}, " ")
}
