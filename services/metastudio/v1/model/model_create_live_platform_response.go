package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLivePlatformResponse Response Object
type CreateLivePlatformResponse struct {

	// 平台ID
	PlatformId *string `json:"platform_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateLivePlatformResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLivePlatformResponse struct{}"
	}

	return strings.Join([]string{"CreateLivePlatformResponse", string(data)}, " ")
}
