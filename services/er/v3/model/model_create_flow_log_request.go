package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFlowLogRequest Request Object
type CreateFlowLogRequest struct {

	// 幂等性标识
	XClientToken *string `json:"X-Client-Token,omitempty"`

	// 企业路由器实例ID
	ErId string `json:"er_id"`

	Body *CreateFlowLogRequestBody `json:"body,omitempty"`
}

func (o CreateFlowLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlowLogRequest struct{}"
	}

	return strings.Join([]string{"CreateFlowLogRequest", string(data)}, " ")
}
