package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTaskResponse Response Object
type UpdateTaskResponse struct {

	// **参数解释**: 错误码，当请求成功时值为 \"00000000\" **取值范围**: 不涉及
	Code *string `json:"code,omitempty"`

	// **参数解释**: 错误信息的描述 **取值范围**: 不涉及
	Message *string `json:"message,omitempty"`

	// **参数解释**: 请求的ID **约束限制**: 不涉及
	RequestId *string `json:"request_id,omitempty"`

	// **参数解释**: 是否成功 **取值范围**: - true  成功 - false 失败
	Success *bool `json:"success,omitempty"`

	Data           *TaskInfo `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o UpdateTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskResponse struct{}"
	}

	return strings.Join([]string{"UpdateTaskResponse", string(data)}, " ")
}
