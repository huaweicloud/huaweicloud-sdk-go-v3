package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMfaResponse Response Object
type ListMfaResponse struct {

	// mfa方式个数
	Count *int32 `json:"count,omitempty"`

	// mfa方式列表
	Methods        *[]MfaRsp `json:"methods,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListMfaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMfaResponse struct{}"
	}

	return strings.Join([]string{"ListMfaResponse", string(data)}, " ")
}
