package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AdOuInfo struct {

	// 名称
	OuName *string `json:"ou_name,omitempty"`

	// 域名地址
	OuDn *string `json:"ou_dn,omitempty"`
}

func (o AdOuInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdOuInfo struct{}"
	}

	return strings.Join([]string{"AdOuInfo", string(data)}, " ")
}
