package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateGlobalEipTagsResponse Response Object
type BatchCreateGlobalEipTagsResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchCreateGlobalEipTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateGlobalEipTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateGlobalEipTagsResponse", string(data)}, " ")
}
