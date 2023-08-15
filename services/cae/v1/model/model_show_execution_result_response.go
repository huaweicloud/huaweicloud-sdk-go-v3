package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowExecutionResultResponse Response Object
type ShowExecutionResultResponse struct {

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion *string `json:"api_version,omitempty"`

	// API类型，固定值“TimerRule”，该值不可修改。
	Kind *string `json:"kind,omitempty"`

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
