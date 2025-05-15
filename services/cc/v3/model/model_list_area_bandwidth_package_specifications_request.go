package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAreaBandwidthPackageSpecificationsRequest Request Object
type ListAreaBandwidthPackageSpecificationsRequest struct {

	// 根据本端大区ID过滤带宽包资源规格列表
	LocalAreaId *[]string `json:"local_area_id,omitempty"`

	// 根据对端大区ID过滤带宽包资源规格列表
	RemoteAreaId *[]string `json:"remote_area_id,omitempty"`
}

func (o ListAreaBandwidthPackageSpecificationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAreaBandwidthPackageSpecificationsRequest struct{}"
	}

	return strings.Join([]string{"ListAreaBandwidthPackageSpecificationsRequest", string(data)}, " ")
}
