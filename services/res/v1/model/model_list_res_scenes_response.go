package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResScenesResponse Response Object
type ListResScenesResponse struct {

	// 是否成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 场景列表信息。
	ResScenes *[]ResScene `json:"res_scenes,omitempty"`

	// 返回消息（请求成功时，不返回此字段）。
	Message *string `json:"message,omitempty"`

	// 错误码（请求成功时，不返回此字段）。
	ErrorCode      *string `json:"error_code,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListResScenesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResScenesResponse struct{}"
	}

	return strings.Join([]string{"ListResScenesResponse", string(data)}, " ")
}
