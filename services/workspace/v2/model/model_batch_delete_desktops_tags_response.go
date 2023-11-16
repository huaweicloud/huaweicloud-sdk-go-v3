package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteDesktopsTagsResponse Response Object
type BatchDeleteDesktopsTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteDesktopsTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteDesktopsTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteDesktopsTagsResponse", string(data)}, " ")
}
