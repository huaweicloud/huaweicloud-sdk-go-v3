package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFunctionAppRequest Request Object
type CreateFunctionAppRequest struct {
	Body *CreateFunctionAppRequestBody `json:"body,omitempty"`
}

func (o CreateFunctionAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFunctionAppRequest struct{}"
	}

	return strings.Join([]string{"CreateFunctionAppRequest", string(data)}, " ")
}
