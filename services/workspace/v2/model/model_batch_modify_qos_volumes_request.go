package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchModifyQosVolumesRequest Request Object
type BatchModifyQosVolumesRequest struct {
	Body *BatchModifyQosVolumesReq `json:"body,omitempty"`
}

func (o BatchModifyQosVolumesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchModifyQosVolumesRequest struct{}"
	}

	return strings.Join([]string{"BatchModifyQosVolumesRequest", string(data)}, " ")
}
