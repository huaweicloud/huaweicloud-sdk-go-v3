package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLayoutResponse Response Object
type ShowLayoutResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 响应信息
	Message *string `json:"message,omitempty"`

	Data *LayoutDetailInfo `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowLayoutResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLayoutResponse struct{}"
	}

	return strings.Join([]string{"ShowLayoutResponse", string(data)}, " ")
}
