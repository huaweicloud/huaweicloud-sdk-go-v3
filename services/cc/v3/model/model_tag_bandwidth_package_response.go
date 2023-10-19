package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagBandwidthPackageResponse Response Object
type TagBandwidthPackageResponse struct {
	XRequestId     *string `json:"x-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o TagBandwidthPackageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagBandwidthPackageResponse struct{}"
	}

	return strings.Join([]string{"TagBandwidthPackageResponse", string(data)}, " ")
}
