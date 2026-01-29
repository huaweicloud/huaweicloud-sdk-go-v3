package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteClassifierResponse Response Object
type DeleteClassifierResponse struct {

	// **参数解释**: 错误码 **取值范围**: 不涉及
	Code *string `json:"code,omitempty"`

	// **参数解释**: 错误描述 **取值范围**: 不涉及
	Message *string `json:"message,omitempty"`

	Data           *DeleteDpeDetail `json:"data,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o DeleteClassifierResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClassifierResponse struct{}"
	}

	return strings.Join([]string{"DeleteClassifierResponse", string(data)}, " ")
}
