package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreatePublicZonesTaskRequest Request Object
type BatchCreatePublicZonesTaskRequest struct {
	Body *BatchCreatePublicZonesTaskRequestBody `json:"body,omitempty"`
}

func (o BatchCreatePublicZonesTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreatePublicZonesTaskRequest struct{}"
	}

	return strings.Join([]string{"BatchCreatePublicZonesTaskRequest", string(data)}, " ")
}
