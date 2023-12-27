package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFacialAnimationsRequest Request Object
type CreateFacialAnimationsRequest struct {
	Body *CreateFasReq `json:"body,omitempty"`
}

func (o CreateFacialAnimationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFacialAnimationsRequest struct{}"
	}

	return strings.Join([]string{"CreateFacialAnimationsRequest", string(data)}, " ")
}
