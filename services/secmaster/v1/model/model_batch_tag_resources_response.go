package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchTagResourcesResponse Response Object
type BatchTagResourcesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchTagResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchTagResourcesResponse struct{}"
	}

	return strings.Join([]string{"BatchTagResourcesResponse", string(data)}, " ")
}
