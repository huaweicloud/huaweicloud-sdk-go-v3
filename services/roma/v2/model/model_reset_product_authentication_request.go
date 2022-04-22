package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ResetProductAuthenticationRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 产品ID
	ProductId int32 `json:"product_id"`
}

func (o ResetProductAuthenticationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetProductAuthenticationRequest struct{}"
	}

	return strings.Join([]string{"ResetProductAuthenticationRequest", string(data)}, " ")
}
