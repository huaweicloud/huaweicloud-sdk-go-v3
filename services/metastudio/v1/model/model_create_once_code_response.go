package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOnceCodeResponse Response Object
type CreateOnceCodeResponse struct {

	// 一次性鉴权码。
	OnceCode *string `json:"once_code,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateOnceCodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOnceCodeResponse struct{}"
	}

	return strings.Join([]string{"CreateOnceCodeResponse", string(data)}, " ")
}
