package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAuthInfoRequest Request Object
type CreateAuthInfoRequest struct {
	Body *CreateAuthInfoReq `json:"body,omitempty"`
}

func (o CreateAuthInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAuthInfoRequest struct{}"
	}

	return strings.Join([]string{"CreateAuthInfoRequest", string(data)}, " ")
}
