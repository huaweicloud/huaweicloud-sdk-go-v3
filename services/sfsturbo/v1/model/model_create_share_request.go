package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateShareRequest struct {
	// MIME类型

	ContentType string `json:"Content-Type"`

	Body *CreateShareRequestBody `json:"body,omitempty"`
}

func (o CreateShareRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateShareRequest struct{}"
	}

	return strings.Join([]string{"CreateShareRequest", string(data)}, " ")
}
