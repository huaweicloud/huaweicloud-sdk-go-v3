package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAccessPointRequest Request Object
type CreateAccessPointRequest struct {
	Body *CreateAccessPointRequestBody `json:"body,omitempty"`
}

func (o CreateAccessPointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAccessPointRequest struct{}"
	}

	return strings.Join([]string{"CreateAccessPointRequest", string(data)}, " ")
}
