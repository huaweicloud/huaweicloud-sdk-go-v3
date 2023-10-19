package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UntagBandwidthPackageResponse Response Object
type UntagBandwidthPackageResponse struct {
	XRequestId     *string `json:"x-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UntagBandwidthPackageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UntagBandwidthPackageResponse struct{}"
	}

	return strings.Join([]string{"UntagBandwidthPackageResponse", string(data)}, " ")
}
