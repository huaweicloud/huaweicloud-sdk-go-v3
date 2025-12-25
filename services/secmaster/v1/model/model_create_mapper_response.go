package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMapperResponse Response Object
type CreateMapperResponse struct {

	// **参数解释**: 错误码 **取值范围**: 不涉及
	Code *string `json:"code,omitempty"`

	// **参数解释**: 错误描述 **取值范围**: 不涉及
	Message *string `json:"message,omitempty"`

	Data           *DpeMappingDetail `json:"data,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o CreateMapperResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMapperResponse struct{}"
	}

	return strings.Join([]string{"CreateMapperResponse", string(data)}, " ")
}
