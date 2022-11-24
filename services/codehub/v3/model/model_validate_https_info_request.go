package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ValidateHttpsInfoRequest struct {

	// 用户iam_user_uuid
	IamUserUuid string `json:"iam_user_uuid"`

	Body *PasswordRequest `json:"body,omitempty"`
}

func (o ValidateHttpsInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateHttpsInfoRequest struct{}"
	}

	return strings.Join([]string{"ValidateHttpsInfoRequest", string(data)}, " ")
}
