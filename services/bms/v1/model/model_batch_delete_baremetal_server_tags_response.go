package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteBaremetalServerTagsResponse Response Object
type BatchDeleteBaremetalServerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteBaremetalServerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteBaremetalServerTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteBaremetalServerTagsResponse", string(data)}, " ")
}
