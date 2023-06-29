package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateApplyRequest Request Object
type CreateApplyRequest struct {
	Body *JoinRequestSchema `json:"body,omitempty"`
}

func (o CreateApplyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApplyRequest struct{}"
	}

	return strings.Join([]string{"CreateApplyRequest", string(data)}, " ")
}
