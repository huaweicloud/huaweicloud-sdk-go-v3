package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateUserQuotasResponse Response Object
type CreateUserQuotasResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateUserQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUserQuotasResponse struct{}"
	}

	return strings.Join([]string{"CreateUserQuotasResponse", string(data)}, " ")
}
