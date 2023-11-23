package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRemoveAvailableZonesRequestBody 移除负载均衡器可用区请求body。
type BatchRemoveAvailableZonesRequestBody struct {

	// 移除的可用区列表，不能为空。
	AvailabilityZoneList []string `json:"availability_zone_list"`
}

func (o BatchRemoveAvailableZonesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRemoveAvailableZonesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchRemoveAvailableZonesRequestBody", string(data)}, " ")
}
