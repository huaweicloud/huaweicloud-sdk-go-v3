package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateCatalogueResponse Response Object
type BatchUpdateCatalogueResponse struct {

	// **参数解释**: 错误码 **取值范围**: 不涉及
	Code *string `json:"code,omitempty"`

	// **参数解释**: 错误描述 **取值范围**: 不涉及
	Message *string `json:"message,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchUpdateCatalogueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateCatalogueResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateCatalogueResponse", string(data)}, " ")
}
