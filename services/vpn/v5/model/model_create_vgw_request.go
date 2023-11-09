package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVgwRequest Request Object
type CreateVgwRequest struct {
	Body *CreateVgwRequestBody `json:"body,omitempty"`
}

func (o CreateVgwRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVgwRequest struct{}"
	}

	return strings.Join([]string{"CreateVgwRequest", string(data)}, " ")
}
