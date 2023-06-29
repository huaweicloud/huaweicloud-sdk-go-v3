package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableSimCardRequest Request Object
type EnableSimCardRequest struct {

	// SIM卡标识，可通过[查询SIM卡列表接口](https://support.huaweicloud.com/api-ocgsl/gsl_07_0008.html)获取
	SimCardId int64 `json:"sim_card_id"`
}

func (o EnableSimCardRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableSimCardRequest struct{}"
	}

	return strings.Join([]string{"EnableSimCardRequest", string(data)}, " ")
}
