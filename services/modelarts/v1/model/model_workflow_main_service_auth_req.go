package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkflowMainServiceAuthReq workflow main service auth request
type WorkflowMainServiceAuthReq struct {

	// 在线服务ID。
	MainServiceId *string `json:"main_service_id,omitempty"`

	// Gallery资产ID。
	ContentId *string `json:"content_id,omitempty"`

	Consume *WorkflowConsume `json:"consume,omitempty"`
}

func (o WorkflowMainServiceAuthReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowMainServiceAuthReq struct{}"
	}

	return strings.Join([]string{"WorkflowMainServiceAuthReq", string(data)}, " ")
}
