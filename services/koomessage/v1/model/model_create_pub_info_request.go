package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePubInfoRequest Request Object
type CreatePubInfoRequest struct {
	Body *CreatePubInfoRequestBody `json:"body,omitempty"`
}

func (o CreatePubInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePubInfoRequest struct{}"
	}

	return strings.Join([]string{"CreatePubInfoRequest", string(data)}, " ")
}
