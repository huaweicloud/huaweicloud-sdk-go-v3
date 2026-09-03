package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTestsuiteInfoUsingResponse Response Object
type UpdateTestsuiteInfoUsingResponse struct {
	Error *Error `json:"error,omitempty"`

	// 对应traceId
	EtTraceId *string `json:"et_trace_id,omitempty"`

	// 响应结果
	Result *interface{} `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateTestsuiteInfoUsingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTestsuiteInfoUsingResponse struct{}"
	}

	return strings.Join([]string{"UpdateTestsuiteInfoUsingResponse", string(data)}, " ")
}
