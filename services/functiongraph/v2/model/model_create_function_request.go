package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFunctionRequest Request Object
type CreateFunctionRequest struct {
	Body *CreateFunctionRequestBody `json:"body,omitempty"`
}

func (o CreateFunctionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFunctionRequest struct{}"
	}

	return strings.Join([]string{"CreateFunctionRequest", string(data)}, " ")
}
