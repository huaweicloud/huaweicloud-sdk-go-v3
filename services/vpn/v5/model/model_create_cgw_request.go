package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCgwRequest Request Object
type CreateCgwRequest struct {
	Body *CreateCgwRequestBody `json:"body,omitempty"`
}

func (o CreateCgwRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCgwRequest struct{}"
	}

	return strings.Join([]string{"CreateCgwRequest", string(data)}, " ")
}
