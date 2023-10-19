package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBandwidthPackageRequest Request Object
type DeleteBandwidthPackageRequest struct {

	// 资源的Id。
	Id string `json:"id"`
}

func (o DeleteBandwidthPackageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBandwidthPackageRequest struct{}"
	}

	return strings.Join([]string{"DeleteBandwidthPackageRequest", string(data)}, " ")
}
