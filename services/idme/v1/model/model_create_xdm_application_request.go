package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateXdmApplicationRequest struct {
	Body *CreateXdmApplicationRequestBody `json:"body,omitempty"`
}

func (o CreateXdmApplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateXdmApplicationRequest struct{}"
	}

	return strings.Join([]string{"CreateXdmApplicationRequest", string(data)}, " ")
}
