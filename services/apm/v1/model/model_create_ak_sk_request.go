package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAkSkRequest Request Object
type CreateAkSkRequest struct {
	Body *CreateAkskModel `json:"body,omitempty"`
}

func (o CreateAkSkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAkSkRequest struct{}"
	}

	return strings.Join([]string{"CreateAkSkRequest", string(data)}, " ")
}
