package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OcaIpVendor 供应商
type OcaIpVendor struct {

	// 供应名称
	Name *string `json:"name,omitempty"`

	// 供应商是否是国产
	IsXc *bool `json:"is_xc,omitempty"`
}

func (o OcaIpVendor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OcaIpVendor struct{}"
	}

	return strings.Join([]string{"OcaIpVendor", string(data)}, " ")
}
