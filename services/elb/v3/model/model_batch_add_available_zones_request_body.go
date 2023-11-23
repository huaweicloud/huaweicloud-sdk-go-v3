package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddAvailableZonesRequestBody 新增负载均衡器可用区请求体
type BatchAddAvailableZonesRequestBody struct {

	// 新增的可用区列表，不能为空。
	AvailabilityZoneList []string `json:"availability_zone_list"`
}

func (o BatchAddAvailableZonesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddAvailableZonesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchAddAvailableZonesRequestBody", string(data)}, " ")
}
