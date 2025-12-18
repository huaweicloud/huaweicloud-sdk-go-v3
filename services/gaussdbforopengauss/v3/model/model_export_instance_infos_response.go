package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportInstanceInfosResponse Response Object
type ExportInstanceInfosResponse struct {

	// **参数解释**: csv格式字符串 **取值范围**: 不涉及
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportInstanceInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportInstanceInfosResponse struct{}"
	}

	return strings.Join([]string{"ExportInstanceInfosResponse", string(data)}, " ")
}
