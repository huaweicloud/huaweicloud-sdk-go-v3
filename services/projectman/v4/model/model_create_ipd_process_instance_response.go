package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIpdProcessInstanceResponse Response Object
type CreateIpdProcessInstanceResponse struct {

	// **参数解释**： 返回状态。 **取值范围**： success：响应成功。 error：响应失败
	Status *string `json:"status,omitempty"`

	// 返回消息
	Message *string `json:"message,omitempty"`

	Result         *ProcessInstanceResponseResult `json:"result,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o CreateIpdProcessInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpdProcessInstanceResponse struct{}"
	}

	return strings.Join([]string{"CreateIpdProcessInstanceResponse", string(data)}, " ")
}
