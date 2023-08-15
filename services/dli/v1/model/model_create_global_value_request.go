package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGlobalValueRequest Request Object
type CreateGlobalValueRequest struct {
	Body *CreateGlobalValueReq `json:"body,omitempty"`
}

func (o CreateGlobalValueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGlobalValueRequest struct{}"
	}

	return strings.Join([]string{"CreateGlobalValueRequest", string(data)}, " ")
}
