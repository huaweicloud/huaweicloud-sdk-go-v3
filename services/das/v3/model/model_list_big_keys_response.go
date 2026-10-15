package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBigKeysResponse Response Object
type ListBigKeysResponse struct {

	// 请求是否成功
	Success *bool `json:"success,omitempty"`

	// 错误码，请求成功时为空
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述，请求成功时为空
	ErrorMsg *string `json:"error_msg,omitempty"`

	Data           *BigKeyListData `json:"data,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListBigKeysResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBigKeysResponse struct{}"
	}

	return strings.Join([]string{"ListBigKeysResponse", string(data)}, " ")
}
