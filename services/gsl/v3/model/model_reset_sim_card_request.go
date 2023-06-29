package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetSimCardRequest Request Object
type ResetSimCardRequest struct {

	// SIM卡标识，可通过[查询SIM卡列表接口](https://support.huaweicloud.com/api-ocgsl/gsl_07_0008.html)获取
	SimCardId int64 `json:"sim_card_id"`

	Body *DownUpTimeForSimCardReq `json:"body,omitempty"`
}

func (o ResetSimCardRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetSimCardRequest struct{}"
	}

	return strings.Join([]string{"ResetSimCardRequest", string(data)}, " ")
}
