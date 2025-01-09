package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHourPackagesTypeResponse Response Object
type ListHourPackagesTypeResponse struct {

	// 可订购小时包类型列表。
	HourPackages *[]HourPackage `json:"hour_packages,omitempty"`

	// 云桌面支持的可用分区列表总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListHourPackagesTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHourPackagesTypeResponse struct{}"
	}

	return strings.Join([]string{"ListHourPackagesTypeResponse", string(data)}, " ")
}
