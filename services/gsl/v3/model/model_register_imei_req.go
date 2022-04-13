package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RegisterImeiReq struct {
	// 绑定类型(1:普通机卡重绑，2：固定机卡重绑)

	BindType int32 `json:"bind_type"`
	// 设备IMEI,84584xxxxxx

	Imei *string `json:"imei,omitempty"`
}

func (o RegisterImeiReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterImeiReq struct{}"
	}

	return strings.Join([]string{"RegisterImeiReq", string(data)}, " ")
}
