package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchResetPasswordResponse Response Object
type BatchResetPasswordResponse struct {
	Results *[]ModifyDbPwdResp `json:"results,omitempty"`

	// 总数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchResetPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchResetPasswordResponse struct{}"
	}

	return strings.Join([]string{"BatchResetPasswordResponse", string(data)}, " ")
}
