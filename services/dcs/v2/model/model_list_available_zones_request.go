package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAvailableZonesRequest Request Object
type ListAvailableZonesRequest struct {
}

func (o ListAvailableZonesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableZonesRequest struct{}"
	}

	return strings.Join([]string{"ListAvailableZonesRequest", string(data)}, " ")
}
