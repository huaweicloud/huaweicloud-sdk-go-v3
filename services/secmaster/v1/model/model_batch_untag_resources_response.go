package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUntagResourcesResponse Response Object
type BatchUntagResourcesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchUntagResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUntagResourcesResponse struct{}"
	}

	return strings.Join([]string{"BatchUntagResourcesResponse", string(data)}, " ")
}
