package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAopWorkflowVersionResponse Response Object
type UpdateAopWorkflowVersionResponse struct {

	// **参数解释**: 错误码 **取值范围**: 不涉及
	Code *string `json:"code,omitempty"`

	// **参数解释**: 错误描述 **取值范围**: 不涉及
	Message *string `json:"message,omitempty"`

	// **参数解释**: 是否成功 **取值范围**: - true  成功 - false 失败
	Success *bool `json:"success,omitempty"`

	// **参数解释**: 请求的ID **约束限制**: 不涉及
	RequestId *string `json:"request_id,omitempty"`

	Data           *AopWorkflowVersionInfo `json:"data,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o UpdateAopWorkflowVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAopWorkflowVersionResponse struct{}"
	}

	return strings.Join([]string{"UpdateAopWorkflowVersionResponse", string(data)}, " ")
}
