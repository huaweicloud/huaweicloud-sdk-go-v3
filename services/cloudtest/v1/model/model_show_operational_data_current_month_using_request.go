package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOperationalDataCurrentMonthUsingRequest Request Object
type ShowOperationalDataCurrentMonthUsingRequest struct {

	// 服务id
	ServiceId string `json:"service_id"`
}

func (o ShowOperationalDataCurrentMonthUsingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOperationalDataCurrentMonthUsingRequest struct{}"
	}

	return strings.Join([]string{"ShowOperationalDataCurrentMonthUsingRequest", string(data)}, " ")
}
