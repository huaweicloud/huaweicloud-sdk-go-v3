package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIndicatorResponse Response Object
type DeleteIndicatorResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	Data *IndicatorBatchOperateResponse `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteIndicatorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIndicatorResponse struct{}"
	}

	return strings.Join([]string{"DeleteIndicatorResponse", string(data)}, " ")
}
