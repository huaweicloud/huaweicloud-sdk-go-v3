package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UntagCloudConnectionResponse Response Object
type UntagCloudConnectionResponse struct {
	XRequestId     *string `json:"x-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UntagCloudConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UntagCloudConnectionResponse struct{}"
	}

	return strings.Join([]string{"UntagCloudConnectionResponse", string(data)}, " ")
}
