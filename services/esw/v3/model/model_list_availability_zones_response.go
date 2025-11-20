package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAvailabilityZonesResponse Response Object
type ListAvailabilityZonesResponse struct {

	// - 参数解释：查询可用区的响应体。 - 约束限制：不涉及。 - 取值范围：不涉及。 - 默认取值：不涉及。
	AvailabilityZones *[]string `json:"availability_zones,omitempty"`

	// - 参数解释：请求的唯一标识。 - 约束限制：UUID格式。 - 取值范围：不涉及。 - 默认取值：不涉及。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListAvailabilityZonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailabilityZonesResponse struct{}"
	}

	return strings.Join([]string{"ListAvailabilityZonesResponse", string(data)}, " ")
}
