package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUserDisclaimerResponse Response Object
type ShowUserDisclaimerResponse struct {

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty"`

	UserDisclaimerRecord *ShowUserDisclaimerRecord `json:"user_disclaimer_record,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowUserDisclaimerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserDisclaimerResponse struct{}"
	}

	return strings.Join([]string{"ShowUserDisclaimerResponse", string(data)}, " ")
}
