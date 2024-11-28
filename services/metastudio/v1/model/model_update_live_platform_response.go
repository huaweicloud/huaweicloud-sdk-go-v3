package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLivePlatformResponse Response Object
type UpdateLivePlatformResponse struct {

	// 平台ID
	PlatformId *string `json:"platform_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateLivePlatformResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLivePlatformResponse struct{}"
	}

	return strings.Join([]string{"UpdateLivePlatformResponse", string(data)}, " ")
}
