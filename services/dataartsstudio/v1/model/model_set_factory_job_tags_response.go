package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetFactoryJobTagsResponse Response Object
type SetFactoryJobTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetFactoryJobTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetFactoryJobTagsResponse struct{}"
	}

	return strings.Join([]string{"SetFactoryJobTagsResponse", string(data)}, " ")
}
