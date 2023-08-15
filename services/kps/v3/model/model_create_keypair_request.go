package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateKeypairRequest Request Object
type CreateKeypairRequest struct {
	Body *CreateKeypairRequestBody `json:"body,omitempty"`
}

func (o CreateKeypairRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKeypairRequest struct{}"
	}

	return strings.Join([]string{"CreateKeypairRequest", string(data)}, " ")
}
