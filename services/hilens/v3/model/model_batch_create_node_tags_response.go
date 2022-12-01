package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchCreateNodeTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateNodeTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateNodeTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateNodeTagsResponse", string(data)}, " ")
}
