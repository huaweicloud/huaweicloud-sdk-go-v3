package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSlowSqlStackResponse Response Object
type ShowSlowSqlStackResponse struct {

	// **参数解释**: 堆栈信息。 **取值范围**: 不涉及。
	GsStack        *string `json:"gs_stack,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSlowSqlStackResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSlowSqlStackResponse struct{}"
	}

	return strings.Join([]string{"ShowSlowSqlStackResponse", string(data)}, " ")
}
