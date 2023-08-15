package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateCodeRequest Request Object
type ValidateCodeRequest struct {

	// 用户id
	UserId string `json:"user_id"`

	Body *CodeVerifyReq `json:"body,omitempty"`
}

func (o ValidateCodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateCodeRequest struct{}"
	}

	return strings.Join([]string{"ValidateCodeRequest", string(data)}, " ")
}
