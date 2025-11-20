package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteUserQuotasResponse Response Object
type DeleteUserQuotasResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteUserQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteUserQuotasResponse struct{}"
	}

	return strings.Join([]string{"DeleteUserQuotasResponse", string(data)}, " ")
}
