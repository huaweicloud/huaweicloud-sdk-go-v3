package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteZonesRequest Request Object
type BatchDeleteZonesRequest struct {
	Body *BatchDeleteZonesRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteZonesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteZonesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteZonesRequest", string(data)}, " ")
}
