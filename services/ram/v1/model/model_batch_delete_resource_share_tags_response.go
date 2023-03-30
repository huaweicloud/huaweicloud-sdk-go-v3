package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchDeleteResourceShareTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteResourceShareTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteResourceShareTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteResourceShareTagsResponse", string(data)}, " ")
}
