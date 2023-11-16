package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePackageIpResponse Response Object
type UpdatePackageIpResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdatePackageIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePackageIpResponse struct{}"
	}

	return strings.Join([]string{"UpdatePackageIpResponse", string(data)}, " ")
}
