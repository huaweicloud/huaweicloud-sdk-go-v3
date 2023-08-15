package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePostalRequest Request Object
type DeletePostalRequest struct {

	// 邮寄地址ID
	AddressId string `json:"address_id"`
}

func (o DeletePostalRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePostalRequest struct{}"
	}

	return strings.Join([]string{"DeletePostalRequest", string(data)}, " ")
}
