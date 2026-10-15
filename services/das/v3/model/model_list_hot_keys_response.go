package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHotKeysResponse Response Object
type ListHotKeysResponse struct {

	// 请求是否成功
	Success *bool `json:"success,omitempty"`

	// 错误码，请求成功时为空
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述，请求成功时为空
	ErrorMsg *string `json:"error_msg,omitempty"`

	Data           *HotKeyListData `json:"data,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListHotKeysResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHotKeysResponse struct{}"
	}

	return strings.Join([]string{"ListHotKeysResponse", string(data)}, " ")
}
