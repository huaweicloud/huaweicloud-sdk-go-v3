package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetSpeedValueRequest Request Object
type SetSpeedValueRequest struct {

	// SIM卡标识，可通过[查询SIM卡列表接口](https://support.huaweicloud.com/api-ocgsl/gsl_07_0008.html)获取
	SimCardId int64 `json:"sim_card_id"`

	Body *SetSpeedValueReq `json:"body,omitempty"`
}

func (o SetSpeedValueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetSpeedValueRequest struct{}"
	}

	return strings.Join([]string{"SetSpeedValueRequest", string(data)}, " ")
}
