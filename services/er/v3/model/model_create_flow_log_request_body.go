package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFlowLogRequestBody 创建流日志请求体
type CreateFlowLogRequestBody struct {
	FlowLog *FlowLogRequest `json:"flow_log"`
}

func (o CreateFlowLogRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlowLogRequestBody struct{}"
	}

	return strings.Join([]string{"CreateFlowLogRequestBody", string(data)}, " ")
}
