package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSshKeysResponse struct {
	Error *Error `json:"error,omitempty"`

	Result *PublicKeyList `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSshKeysResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSshKeysResponse struct{}"
	}

	return strings.Join([]string{"ListSshKeysResponse", string(data)}, " ")
}
