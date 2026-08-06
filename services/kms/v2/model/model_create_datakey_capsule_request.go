package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDatakeyCapsuleRequest Request Object
type CreateDatakeyCapsuleRequest struct {
	Body *CreateDatakeyCapsuleRequestBody `json:"body,omitempty"`
}

func (o CreateDatakeyCapsuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatakeyCapsuleRequest struct{}"
	}

	return strings.Join([]string{"CreateDatakeyCapsuleRequest", string(data)}, " ")
}
