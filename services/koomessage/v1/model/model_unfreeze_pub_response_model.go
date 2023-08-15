package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UnfreezePubResponseModel struct {

	// 服务号ID。
	PubId *string `json:"pub_id,omitempty"`
}

func (o UnfreezePubResponseModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnfreezePubResponseModel struct{}"
	}

	return strings.Join([]string{"UnfreezePubResponseModel", string(data)}, " ")
}
