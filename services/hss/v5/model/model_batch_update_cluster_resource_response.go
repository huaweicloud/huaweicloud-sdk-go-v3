package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateClusterResourceResponse Response Object
type BatchUpdateClusterResourceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchUpdateClusterResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateClusterResourceResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateClusterResourceResponse", string(data)}, " ")
}
