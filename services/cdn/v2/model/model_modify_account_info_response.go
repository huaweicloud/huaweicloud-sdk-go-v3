package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyAccountInfoResponse Response Object
type ModifyAccountInfoResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ModifyAccountInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyAccountInfoResponse struct{}"
	}

	return strings.Join([]string{"ModifyAccountInfoResponse", string(data)}, " ")
}
