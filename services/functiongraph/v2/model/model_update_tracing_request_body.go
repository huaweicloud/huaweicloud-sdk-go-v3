package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTracingRequestBody 更新函数调用链请求体
type UpdateTracingRequestBody struct {

	// apm的ak
	TracingAk *string `json:"tracing_ak,omitempty"`

	// apm的sk
	TracingSk *string `json:"tracing_sk,omitempty"`
}

func (o UpdateTracingRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTracingRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateTracingRequestBody", string(data)}, " ")
}
