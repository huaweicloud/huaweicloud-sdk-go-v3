package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGlobalEipRequest Request Object
type CreateGlobalEipRequest struct {
	Body *CreateGlobalEipRequestBody `json:"body,omitempty"`
}

func (o CreateGlobalEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGlobalEipRequest struct{}"
	}

	return strings.Join([]string{"CreateGlobalEipRequest", string(data)}, " ")
}
