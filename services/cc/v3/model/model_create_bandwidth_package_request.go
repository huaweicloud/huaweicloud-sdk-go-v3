package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBandwidthPackageRequest Request Object
type CreateBandwidthPackageRequest struct {
	Body *CreateBandwidthPackageRequestBody `json:"body,omitempty"`
}

func (o CreateBandwidthPackageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBandwidthPackageRequest struct{}"
	}

	return strings.Join([]string{"CreateBandwidthPackageRequest", string(data)}, " ")
}
