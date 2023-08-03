package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FreezePubResponseModel struct {

	// 服务号ID。
	PubId *string `json:"pub_id,omitempty"`
}

func (o FreezePubResponseModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FreezePubResponseModel struct{}"
	}

	return strings.Join([]string{"FreezePubResponseModel", string(data)}, " ")
}
