package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAimSendReportsRequest Request Object
type ListAimSendReportsRequest struct {

	// 请求体参数类型，该字段必须设置为：application/json。
	ContentType string `json:"Content-Type"`

	Body *ListAimSendReportsRequestBody `json:"body,omitempty"`
}

func (o ListAimSendReportsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAimSendReportsRequest struct{}"
	}

	return strings.Join([]string{"ListAimSendReportsRequest", string(data)}, " ")
}
