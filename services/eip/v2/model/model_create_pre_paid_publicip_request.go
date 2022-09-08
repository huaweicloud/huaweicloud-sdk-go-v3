package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreatePrePaidPublicipRequest struct {
	Body *CreatePrePaidPublicipRequestBody `json:"body,omitempty"`
}

func (o CreatePrePaidPublicipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrePaidPublicipRequest struct{}"
	}

	return strings.Join([]string{"CreatePrePaidPublicipRequest", string(data)}, " ")
}
