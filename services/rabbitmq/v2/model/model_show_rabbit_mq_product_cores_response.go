package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRabbitMqProductCoresResponse Response Object
type ShowRabbitMqProductCoresResponse struct {

	// 核数
	CoreNum        *int32 `json:"core_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowRabbitMqProductCoresResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRabbitMqProductCoresResponse struct{}"
	}

	return strings.Join([]string{"ShowRabbitMqProductCoresResponse", string(data)}, " ")
}
