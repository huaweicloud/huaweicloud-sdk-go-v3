package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClassifierResponse Response Object
type UpdateClassifierResponse struct {

	// **参数解释**: 错误码 **取值范围**: 不涉及
	Code *string `json:"code,omitempty"`

	// **参数解释**: 错误描述 **取值范围**: 不涉及
	Message *string `json:"message,omitempty"`

	Data           *CreateDpeClassifyRequestBody `json:"data,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o UpdateClassifierResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClassifierResponse struct{}"
	}

	return strings.Join([]string{"UpdateClassifierResponse", string(data)}, " ")
}
