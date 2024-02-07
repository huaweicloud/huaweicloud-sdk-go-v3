package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteUserDisclaimerResponse Response Object
type DeleteUserDisclaimerResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteUserDisclaimerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteUserDisclaimerResponse struct{}"
	}

	return strings.Join([]string{"DeleteUserDisclaimerResponse", string(data)}, " ")
}
