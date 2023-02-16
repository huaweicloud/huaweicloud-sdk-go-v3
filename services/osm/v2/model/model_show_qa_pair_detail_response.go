package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowQaPairDetailResponse struct {
	Qapair *QaPair `json:"qapair,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowQaPairDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQaPairDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowQaPairDetailResponse", string(data)}, " ")
}
