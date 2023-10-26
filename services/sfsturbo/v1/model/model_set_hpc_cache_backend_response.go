package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetHpcCacheBackendResponse Response Object
type SetHpcCacheBackendResponse struct {
	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetHpcCacheBackendResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetHpcCacheBackendResponse struct{}"
	}

	return strings.Join([]string{"SetHpcCacheBackendResponse", string(data)}, " ")
}
