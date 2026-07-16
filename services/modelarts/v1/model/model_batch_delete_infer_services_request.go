package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteInferServicesRequest Request Object
type BatchDeleteInferServicesRequest struct {
	Body *DeleteServicesRequest `json:"body,omitempty"`
}

func (o BatchDeleteInferServicesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteInferServicesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteInferServicesRequest", string(data)}, " ")
}
