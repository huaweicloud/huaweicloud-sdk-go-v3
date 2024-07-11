package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SetKernelPluginLicenseRequestBody struct {

	// license
	LicenseStr string `json:"license_str"`
}

func (o SetKernelPluginLicenseRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetKernelPluginLicenseRequestBody struct{}"
	}

	return strings.Join([]string{"SetKernelPluginLicenseRequestBody", string(data)}, " ")
}
