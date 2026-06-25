package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePeriodToOnDemandInstantlyResponse Response Object
type UpdatePeriodToOnDemandInstantlyResponse struct {

	// |参数名称：包年包月即时转按需结果| |参数约束以及描述：包年包月即时转按需结果。HTTP 200的时候返回该字段，具体参见ToOndemandServiceResult。|
	ToOndemandServiceResults *[]ToOndemandServiceResult `json:"to_ondemand_service_results,omitempty"`
	HttpStatusCode           int                        `json:"-"`
}

func (o UpdatePeriodToOnDemandInstantlyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePeriodToOnDemandInstantlyResponse struct{}"
	}

	return strings.Join([]string{"UpdatePeriodToOnDemandInstantlyResponse", string(data)}, " ")
}
