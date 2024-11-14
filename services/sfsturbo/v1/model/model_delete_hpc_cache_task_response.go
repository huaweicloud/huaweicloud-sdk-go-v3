package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHpcCacheTaskResponse Response Object
type DeleteHpcCacheTaskResponse struct {
	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteHpcCacheTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHpcCacheTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteHpcCacheTaskResponse", string(data)}, " ")
}
