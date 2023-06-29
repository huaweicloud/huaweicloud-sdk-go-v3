package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFlowLogRequest Request Object
type CreateFlowLogRequest struct {
	Body *CreateFlowLogReqBody `json:"body,omitempty"`
}

func (o CreateFlowLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlowLogRequest struct{}"
	}

	return strings.Join([]string{"CreateFlowLogRequest", string(data)}, " ")
}
