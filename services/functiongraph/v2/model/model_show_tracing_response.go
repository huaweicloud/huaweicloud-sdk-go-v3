package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTracingResponse Response Object
type ShowTracingResponse struct {

	// apm的ak
	TracingAk *string `json:"tracing_ak,omitempty"`

	// apm的sk
	TracingSk      *string `json:"tracing_sk,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTracingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTracingResponse struct{}"
	}

	return strings.Join([]string{"ShowTracingResponse", string(data)}, " ")
}
