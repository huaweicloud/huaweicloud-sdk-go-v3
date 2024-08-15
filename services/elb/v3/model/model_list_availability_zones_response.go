package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAvailabilityZonesResponse Response Object
type ListAvailabilityZonesResponse struct {

	// 参数解释：请求ID。  注：自动生成。
	RequestId *string `json:"request_id,omitempty"`

	// 参数解释：返回创建LB时可使用的可用区集合列表。如：[[az1,az2],[az2,az3]] ,则在创建LB时，只能从其中的一个子列表中选择一个或多个可用区，不能跨列表选择。在上述例子中，不能选择az1和az3。
	AvailabilityZones *[][]AvailabilityZone `json:"availability_zones,omitempty"`

	// 可用区的产品编码，仅边缘场景有效。
	SpecCode       *string `json:"spec_code,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListAvailabilityZonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailabilityZonesResponse struct{}"
	}

	return strings.Join([]string{"ListAvailabilityZonesResponse", string(data)}, " ")
}
