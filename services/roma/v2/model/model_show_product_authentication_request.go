package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProductAuthenticationRequest Request Object
type ShowProductAuthenticationRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 产品ID
	ProductId int32 `json:"product_id"`
}

func (o ShowProductAuthenticationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProductAuthenticationRequest struct{}"
	}

	return strings.Join([]string{"ShowProductAuthenticationRequest", string(data)}, " ")
}
