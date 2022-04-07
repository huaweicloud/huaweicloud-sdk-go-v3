package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DriverLicenseResultStatus struct {
	Name *string `json:"name,omitempty"`
}

func (o DriverLicenseResultStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DriverLicenseResultStatus struct{}"
	}

	return strings.Join([]string{"DriverLicenseResultStatus", string(data)}, " ")
}
