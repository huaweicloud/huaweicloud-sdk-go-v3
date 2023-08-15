package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateResDatasourceResponse Response Object
type UpdateResDatasourceResponse struct {

	// 是否成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 返回消息。
	Message *string `json:"message,omitempty"`

	// 错误码（请求成功时，不返回此字段）。
	ErrorCode      *string `json:"error_code,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateResDatasourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResDatasourceResponse struct{}"
	}

	return strings.Join([]string{"UpdateResDatasourceResponse", string(data)}, " ")
}
