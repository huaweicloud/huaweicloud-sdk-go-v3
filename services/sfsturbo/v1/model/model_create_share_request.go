package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateShareRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type" xml:"Content-Type"`

	Body *CreateShareRequestBody `json:"body,omitempty" xml:"body"`
}

func (o CreateShareRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateShareRequest struct{}"
	}

	return strings.Join([]string{"CreateShareRequest", string(data)}, " ")
}
