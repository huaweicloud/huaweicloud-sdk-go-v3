package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateModelServiceRequest Request Object
type CreateModelServiceRequest struct {
	Body *CreateModelServiceReq `json:"body,omitempty"`
}

func (o CreateModelServiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateModelServiceRequest struct{}"
	}

	return strings.Join([]string{"CreateModelServiceRequest", string(data)}, " ")
}
