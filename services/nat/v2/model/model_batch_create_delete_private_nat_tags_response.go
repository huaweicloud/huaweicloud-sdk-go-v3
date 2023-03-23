package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchCreateDeletePrivateNatTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateDeletePrivateNatTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateDeletePrivateNatTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateDeletePrivateNatTagsResponse", string(data)}, " ")
}
