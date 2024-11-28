package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLivePlatformResponse Response Object
type DeleteLivePlatformResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteLivePlatformResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLivePlatformResponse struct{}"
	}

	return strings.Join([]string{"DeleteLivePlatformResponse", string(data)}, " ")
}
