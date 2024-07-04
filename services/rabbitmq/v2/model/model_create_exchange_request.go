package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateExchangeRequest Request Object
type CreateExchangeRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// Vhost名称
	Vhost string `json:"vhost"`

	Body *CreateExchangeBody `json:"body,omitempty"`
}

func (o CreateExchangeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExchangeRequest struct{}"
	}

	return strings.Join([]string{"CreateExchangeRequest", string(data)}, " ")
}
