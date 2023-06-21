package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 解绑服务号返回体。
type UnfreezePubResponseModel struct {
	Data *FreezePubResponseModelData `json:"data,omitempty"`
}

func (o UnfreezePubResponseModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnfreezePubResponseModel struct{}"
	}

	return strings.Join([]string{"UnfreezePubResponseModel", string(data)}, " ")
}
