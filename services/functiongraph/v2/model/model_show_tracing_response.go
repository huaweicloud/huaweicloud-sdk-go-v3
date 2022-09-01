package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowTracingResponse struct {

	// apm的ak
	TracingAk *string `json:"tracing_ak,omitempty" xml:"tracing_ak"`

	// apm的sk
	TracingSk      *string `json:"tracing_sk,omitempty" xml:"tracing_sk"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTracingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTracingResponse struct{}"
	}

	return strings.Join([]string{"ShowTracingResponse", string(data)}, " ")
}
