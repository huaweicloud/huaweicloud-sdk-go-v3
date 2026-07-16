package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteInferIntranetConnectionsRequest Request Object
type BatchDeleteInferIntranetConnectionsRequest struct {
	Body *IntranetConnectionDeleteRequest `json:"body,omitempty"`
}

func (o BatchDeleteInferIntranetConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteInferIntranetConnectionsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteInferIntranetConnectionsRequest", string(data)}, " ")
}
