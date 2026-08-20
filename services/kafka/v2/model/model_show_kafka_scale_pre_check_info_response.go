package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKafkaScalePreCheckInfoResponse Response Object
type ShowKafkaScalePreCheckInfoResponse struct {

	// **参数解释**： 扩容前置检查信息。
	Body           *[]ShowKafkaScalePreCheckInfoEntity `json:"body,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ShowKafkaScalePreCheckInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKafkaScalePreCheckInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowKafkaScalePreCheckInfoResponse", string(data)}, " ")
}
