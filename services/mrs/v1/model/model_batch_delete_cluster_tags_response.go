package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchDeleteClusterTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteClusterTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteClusterTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteClusterTagsResponse", string(data)}, " ")
}
