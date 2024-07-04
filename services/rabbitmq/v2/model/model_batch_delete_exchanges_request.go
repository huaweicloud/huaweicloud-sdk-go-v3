package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteExchangesRequest Request Object
type BatchDeleteExchangesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// Vhost名称
	Vhost string `json:"vhost"`

	Body *BatchDeleteBody `json:"body,omitempty"`
}

func (o BatchDeleteExchangesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteExchangesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteExchangesRequest", string(data)}, " ")
}
