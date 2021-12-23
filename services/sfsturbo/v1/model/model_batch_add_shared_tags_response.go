package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchAddSharedTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchAddSharedTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddSharedTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchAddSharedTagsResponse", string(data)}, " ")
}
