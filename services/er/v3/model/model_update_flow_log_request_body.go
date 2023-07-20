package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFlowLogRequestBody 更新流日志的名称/描述
type UpdateFlowLogRequestBody struct {

	// 流日志名称
	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`
}

func (o UpdateFlowLogRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFlowLogRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateFlowLogRequestBody", string(data)}, " ")
}
