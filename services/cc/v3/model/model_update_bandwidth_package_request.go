package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBandwidthPackageRequest Request Object
type UpdateBandwidthPackageRequest struct {

	// 资源的Id。
	Id string `json:"id"`

	Body *UpdateBandwidthPackageRequestBody `json:"body,omitempty"`
}

func (o UpdateBandwidthPackageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBandwidthPackageRequest struct{}"
	}

	return strings.Join([]string{"UpdateBandwidthPackageRequest", string(data)}, " ")
}
