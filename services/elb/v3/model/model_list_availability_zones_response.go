package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAvailabilityZonesResponse struct {

	// 请求ID。  注：自动生成。
	RequestId *string `json:"request_id,omitempty"`

	// 返回创建LB时可使用的可用区集合列表。
	AvailabilityZones *[][]AvailabilityZone `json:"availability_zones,omitempty"`
	HttpStatusCode    int                   `json:"-"`
}

func (o ListAvailabilityZonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailabilityZonesResponse struct{}"
	}

	return strings.Join([]string{"ListAvailabilityZonesResponse", string(data)}, " ")
}
