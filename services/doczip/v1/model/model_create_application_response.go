package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateApplicationResponse Response Object
type CreateApplicationResponse struct {

	// **参数解释**： 应用ID。 **取值范围**： [1-256]个字符。
	ApplicationId *string `json:"application_id,omitempty"`

	XTraceId       *string `json:"X-Trace-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateApplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApplicationResponse struct{}"
	}

	return strings.Join([]string{"CreateApplicationResponse", string(data)}, " ")
}
