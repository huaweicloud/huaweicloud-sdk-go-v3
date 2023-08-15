package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateCaTagsResponse Response Object
type BatchCreateCaTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateCaTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateCaTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateCaTagsResponse", string(data)}, " ")
}
