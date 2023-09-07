package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowExecutionResultResponse Response Object
type ShowExecutionResultResponse struct {
	ApiVersion *ApiVersionObj `json:"api_version,omitempty"`

	Kind *TimeRuleKindObj `json:"kind,omitempty"`

	Spec           *ExecutionDetails `json:"spec,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowExecutionResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExecutionResultResponse struct{}"
	}

	return strings.Join([]string{"ShowExecutionResultResponse", string(data)}, " ")
}
