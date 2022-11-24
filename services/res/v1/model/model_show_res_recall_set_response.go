package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowResRecallSetResponse struct {

	// 候选集列表
	ResultSet *[]ResultSet `json:"result_set,omitempty"`

	// 是否成功
	IsSuccess *bool `json:"is_success,omitempty"`

	// 返回消息（请求成功时，不返回此字段）
	Message *string `json:"message,omitempty"`

	// 错误码（请求成功时，不返回此字段）
	ErrorCode      *string `json:"error_code,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowResRecallSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResRecallSetResponse struct{}"
	}

	return strings.Join([]string{"ShowResRecallSetResponse", string(data)}, " ")
}
