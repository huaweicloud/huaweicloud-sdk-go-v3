package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateAccessConfigRequest struct {
	Body *CreateAccessConfigRequestBody `json:"body,omitempty"`
}

func (o CreateAccessConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAccessConfigRequest struct{}"
	}

	return strings.Join([]string{"CreateAccessConfigRequest", string(data)}, " ")
}
