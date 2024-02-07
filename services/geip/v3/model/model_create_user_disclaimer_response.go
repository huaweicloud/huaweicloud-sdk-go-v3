package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateUserDisclaimerResponse Response Object
type CreateUserDisclaimerResponse struct {

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty"`

	UserDisclaimerRecord *CreateUserDisclaimerRecord `json:"user_disclaimer_record,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateUserDisclaimerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUserDisclaimerResponse struct{}"
	}

	return strings.Join([]string{"CreateUserDisclaimerResponse", string(data)}, " ")
}
