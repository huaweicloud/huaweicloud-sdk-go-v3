package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUserChargeTypeRequest Request Object
type ShowUserChargeTypeRequest struct {
}

func (o ShowUserChargeTypeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserChargeTypeRequest struct{}"
	}

	return strings.Join([]string{"ShowUserChargeTypeRequest", string(data)}, " ")
}
