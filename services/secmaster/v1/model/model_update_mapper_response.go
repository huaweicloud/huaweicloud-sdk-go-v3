package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMapperResponse Response Object
type UpdateMapperResponse struct {

	// **参数解释**: 错误码 **取值范围**: 不涉及
	Code *string `json:"code,omitempty"`

	// **参数解释**: 错误描述 **取值范围**: 不涉及
	Message *string `json:"message,omitempty"`

	Data           *CreateDpeMappingRequestBody `json:"data,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o UpdateMapperResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMapperResponse struct{}"
	}

	return strings.Join([]string{"UpdateMapperResponse", string(data)}, " ")
}
