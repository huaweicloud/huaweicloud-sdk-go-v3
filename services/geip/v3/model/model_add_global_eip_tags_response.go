package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddGlobalEipTagsResponse Response Object
type AddGlobalEipTagsResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddGlobalEipTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddGlobalEipTagsResponse struct{}"
	}

	return strings.Join([]string{"AddGlobalEipTagsResponse", string(data)}, " ")
}
