package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCsrRequest Request Object
type CreateCsrRequest struct {
	Body *CreateCsrRequestBody `json:"body,omitempty"`
}

func (o CreateCsrRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCsrRequest struct{}"
	}

	return strings.Join([]string{"CreateCsrRequest", string(data)}, " ")
}
