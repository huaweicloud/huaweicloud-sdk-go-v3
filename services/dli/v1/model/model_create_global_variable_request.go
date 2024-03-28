package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGlobalVariableRequest Request Object
type CreateGlobalVariableRequest struct {
	Body *CreateGlobalVariableRequestBody `json:"body,omitempty"`
}

func (o CreateGlobalVariableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGlobalVariableRequest struct{}"
	}

	return strings.Join([]string{"CreateGlobalVariableRequest", string(data)}, " ")
}
