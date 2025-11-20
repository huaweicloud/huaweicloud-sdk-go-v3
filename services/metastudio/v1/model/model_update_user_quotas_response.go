package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserQuotasResponse Response Object
type UpdateUserQuotasResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateUserQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserQuotasResponse struct{}"
	}

	return strings.Join([]string{"UpdateUserQuotasResponse", string(data)}, " ")
}
