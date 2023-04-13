package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AzCodeInfo struct {

	// 可用区编码。
	AzCode *string `json:"az_code,omitempty"`
}

func (o AzCodeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AzCodeInfo struct{}"
	}

	return strings.Join([]string{"AzCodeInfo", string(data)}, " ")
}
