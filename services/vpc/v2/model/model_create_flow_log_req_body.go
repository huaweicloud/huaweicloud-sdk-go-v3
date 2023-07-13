package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFlowLogReqBody
type CreateFlowLogReqBody struct {
	FlowLog *CreateFlowLogReq `json:"flow_log"`
}

func (o CreateFlowLogReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlowLogReqBody struct{}"
	}

	return strings.Join([]string{"CreateFlowLogReqBody", string(data)}, " ")
}
