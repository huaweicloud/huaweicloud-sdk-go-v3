package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteGlobalEipTagsResponse Response Object
type BatchDeleteGlobalEipTagsResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchDeleteGlobalEipTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteGlobalEipTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteGlobalEipTagsResponse", string(data)}, " ")
}
