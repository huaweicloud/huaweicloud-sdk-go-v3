package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSetZonesStatusRequest Request Object
type BatchSetZonesStatusRequest struct {
	Body *BatchSetZonesStatusRequestBody `json:"body,omitempty"`
}

func (o BatchSetZonesStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetZonesStatusRequest struct{}"
	}

	return strings.Join([]string{"BatchSetZonesStatusRequest", string(data)}, " ")
}
