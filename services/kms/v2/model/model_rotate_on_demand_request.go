package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RotateOnDemandRequest Request Object
type RotateOnDemandRequest struct {
	Body *RotateOnDemandRequestBody `json:"body,omitempty"`
}

func (o RotateOnDemandRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RotateOnDemandRequest struct{}"
	}

	return strings.Join([]string{"RotateOnDemandRequest", string(data)}, " ")
}
