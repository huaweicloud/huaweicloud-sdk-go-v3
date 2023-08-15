package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAttributeRequest Request Object
type CreateAttributeRequest struct {
	Body *AddOrModifyAttributeReq `json:"body,omitempty"`
}

func (o CreateAttributeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAttributeRequest struct{}"
	}

	return strings.Join([]string{"CreateAttributeRequest", string(data)}, " ")
}
