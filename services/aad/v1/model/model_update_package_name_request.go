package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePackageNameRequest Request Object
type UpdatePackageNameRequest struct {

	// 防护包id
	PackageId string `json:"package_id"`

	Body *UpdatePackageNameRequestBody `json:"body,omitempty"`
}

func (o UpdatePackageNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePackageNameRequest struct{}"
	}

	return strings.Join([]string{"UpdatePackageNameRequest", string(data)}, " ")
}
