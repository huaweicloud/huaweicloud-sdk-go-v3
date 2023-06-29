package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProductConfigRequest Request Object
type ShowProductConfigRequest struct {

	// 协议类型
	ProtocolType string `json:"protocol_type"`
}

func (o ShowProductConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProductConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowProductConfigRequest", string(data)}, " ")
}
