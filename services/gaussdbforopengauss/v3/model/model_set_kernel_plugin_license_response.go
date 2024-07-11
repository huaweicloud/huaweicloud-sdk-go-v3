package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetKernelPluginLicenseResponse Response Object
type SetKernelPluginLicenseResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetKernelPluginLicenseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetKernelPluginLicenseResponse struct{}"
	}

	return strings.Join([]string{"SetKernelPluginLicenseResponse", string(data)}, " ")
}
