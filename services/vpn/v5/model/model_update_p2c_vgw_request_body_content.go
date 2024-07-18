package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateP2cVgwRequestBodyContent struct {

	// P2C VPN网关名称。1-64字符，中文、英文、数字包含下划线
	Name *string `json:"name,omitempty"`

	// eip的ID。用于给P2C VPN网关绑定新的EIP，需要先解绑当前的EIP
	EipId *string `json:"eip_id,omitempty"`
}

func (o UpdateP2cVgwRequestBodyContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateP2cVgwRequestBodyContent struct{}"
	}

	return strings.Join([]string{"UpdateP2cVgwRequestBodyContent", string(data)}, " ")
}
