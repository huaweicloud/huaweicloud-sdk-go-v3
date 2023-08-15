package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopSimCardRequest Request Object
type StopSimCardRequest struct {

	// SIM卡标识，可通过[查询SIM卡列表接口](https://support.huaweicloud.com/api-ocgsl/gsl_07_0008.html)获取
	SimCardId int64 `json:"sim_card_id"`

	Body *DownUpTimeForSimCardReq `json:"body,omitempty"`
}

func (o StopSimCardRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopSimCardRequest struct{}"
	}

	return strings.Join([]string{"StopSimCardRequest", string(data)}, " ")
}
