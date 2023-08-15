package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePortResponseModelData 消息体。
type DeletePortResponseModelData struct {

	// 通道号。
	Port string `json:"port"`

	// 通道号类型。 - 1：普通 - 3：前缀号段 - 5：后缀号段
	PortType int32 `json:"port_type"`
}

func (o DeletePortResponseModelData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePortResponseModelData struct{}"
	}

	return strings.Join([]string{"DeletePortResponseModelData", string(data)}, " ")
}
