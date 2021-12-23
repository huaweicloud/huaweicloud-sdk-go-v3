package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ImportFunctionRequest struct {
	Body *ImportFunctionRequestBody `json:"body,omitempty"`
}

func (o ImportFunctionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportFunctionRequest struct{}"
	}

	return strings.Join([]string{"ImportFunctionRequest", string(data)}, " ")
}
