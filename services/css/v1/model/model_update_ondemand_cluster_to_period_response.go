package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOndemandClusterToPeriodResponse Response Object
type UpdateOndemandClusterToPeriodResponse struct {

	// 订单ID（此订单类型为“新购”）。
	OrderId        *string `json:"orderId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateOndemandClusterToPeriodResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOndemandClusterToPeriodResponse struct{}"
	}

	return strings.Join([]string{"UpdateOndemandClusterToPeriodResponse", string(data)}, " ")
}
