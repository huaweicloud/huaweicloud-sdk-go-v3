package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuditResponse Response Object
type ShowAuditResponse struct {

	// 结果状态
	Status *string `json:"status,omitempty"`

	// 请求id
	TraceId *string `json:"trace_id,omitempty"`

	// 请求返回结果，接口不同，返回不同
	Result         *interface{} `json:"result,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowAuditResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuditResponse struct{}"
	}

	return strings.Join([]string{"ShowAuditResponse", string(data)}, " ")
}
