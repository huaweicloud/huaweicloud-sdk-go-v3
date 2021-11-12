package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchCreateOrDeleteResourceTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateOrDeleteResourceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateOrDeleteResourceTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateOrDeleteResourceTagsResponse", string(data)}, " ")
}
