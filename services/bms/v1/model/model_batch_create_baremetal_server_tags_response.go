package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateBaremetalServerTagsResponse Response Object
type BatchCreateBaremetalServerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateBaremetalServerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateBaremetalServerTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateBaremetalServerTagsResponse", string(data)}, " ")
}
