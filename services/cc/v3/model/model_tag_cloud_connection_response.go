package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagCloudConnectionResponse Response Object
type TagCloudConnectionResponse struct {
	XRequestId     *string `json:"x-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o TagCloudConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagCloudConnectionResponse struct{}"
	}

	return strings.Join([]string{"TagCloudConnectionResponse", string(data)}, " ")
}
