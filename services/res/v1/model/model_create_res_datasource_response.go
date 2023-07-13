package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResDatasourceResponse Response Object
type CreateResDatasourceResponse struct {

	// 是否成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	Datasource *Datasource `json:"datasource,omitempty"`

	// 返回消息（请求成功时，不返回此字段）。
	Message *string `json:"message,omitempty"`

	// 错误码（请求成功时，不返回此字段）。
	ErrorCode      *string `json:"error_code,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateResDatasourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResDatasourceResponse struct{}"
	}

	return strings.Join([]string{"CreateResDatasourceResponse", string(data)}, " ")
}
