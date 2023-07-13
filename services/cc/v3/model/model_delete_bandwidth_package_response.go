package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBandwidthPackageResponse Response Object
type DeleteBandwidthPackageResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteBandwidthPackageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBandwidthPackageResponse struct{}"
	}

	return strings.Join([]string{"DeleteBandwidthPackageResponse", string(data)}, " ")
}
