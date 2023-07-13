package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FreezePubResponseModel struct {
	Data *FreezePubResponseModelData `json:"data,omitempty"`
}

func (o FreezePubResponseModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FreezePubResponseModel struct{}"
	}

	return strings.Join([]string{"FreezePubResponseModel", string(data)}, " ")
}
