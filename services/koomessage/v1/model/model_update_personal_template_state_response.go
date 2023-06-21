package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdatePersonalTemplateStateResponse struct {

	// 请求状态，固定200。
	Status *string `json:"status,omitempty"`

	// 状态描述。
	Message *string `json:"message,omitempty"`

	// 固定为null。
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdatePersonalTemplateStateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePersonalTemplateStateResponse struct{}"
	}

	return strings.Join([]string{"UpdatePersonalTemplateStateResponse", string(data)}, " ")
}
