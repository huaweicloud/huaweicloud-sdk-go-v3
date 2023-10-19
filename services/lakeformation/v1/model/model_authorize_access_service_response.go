package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuthorizeAccessServiceResponse Response Object
type AuthorizeAccessServiceResponse struct {
	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AuthorizeAccessServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizeAccessServiceResponse struct{}"
	}

	return strings.Join([]string{"AuthorizeAccessServiceResponse", string(data)}, " ")
}
