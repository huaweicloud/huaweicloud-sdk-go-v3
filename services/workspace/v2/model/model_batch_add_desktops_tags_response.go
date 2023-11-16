package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddDesktopsTagsResponse Response Object
type BatchAddDesktopsTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchAddDesktopsTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddDesktopsTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchAddDesktopsTagsResponse", string(data)}, " ")
}
