package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePtrRequest Request Object
type CreatePtrRequest struct {
	Body *CreatePtrRequestBody `json:"body,omitempty"`
}

func (o CreatePtrRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePtrRequest struct{}"
	}

	return strings.Join([]string{"CreatePtrRequest", string(data)}, " ")
}
