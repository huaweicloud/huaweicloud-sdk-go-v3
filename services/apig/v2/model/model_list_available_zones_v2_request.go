package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAvailableZonesV2Request Request Object
type ListAvailableZonesV2Request struct {
}

func (o ListAvailableZonesV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableZonesV2Request struct{}"
	}

	return strings.Join([]string{"ListAvailableZonesV2Request", string(data)}, " ")
}
