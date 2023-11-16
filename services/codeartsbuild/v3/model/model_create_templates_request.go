package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTemplatesRequest Request Object
type CreateTemplatesRequest struct {
	Body *CreateTemplatesRequestBody `json:"body,omitempty"`
}

func (o CreateTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTemplatesRequest struct{}"
	}

	return strings.Join([]string{"CreateTemplatesRequest", string(data)}, " ")
}
