package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSharesByTagResponse Response Object
type ListSharesByTagResponse struct {
	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSharesByTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSharesByTagResponse struct{}"
	}

	return strings.Join([]string{"ListSharesByTagResponse", string(data)}, " ")
}
