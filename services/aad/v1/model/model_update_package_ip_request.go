package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePackageIpRequest Request Object
type UpdatePackageIpRequest struct {

	// 实例id
	PackageId string `json:"package_id"`

	Body *UpdatePackageIpRequestBody `json:"body,omitempty"`
}

func (o UpdatePackageIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePackageIpRequest struct{}"
	}

	return strings.Join([]string{"UpdatePackageIpRequest", string(data)}, " ")
}
