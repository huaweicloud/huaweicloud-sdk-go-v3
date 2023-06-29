package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FreezePubResponseModelData 请求成功返回的数据。
type FreezePubResponseModelData struct {

	// 服务号ID。
	PubId *string `json:"pub_id,omitempty"`
}

func (o FreezePubResponseModelData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FreezePubResponseModelData struct{}"
	}

	return strings.Join([]string{"FreezePubResponseModelData", string(data)}, " ")
}
