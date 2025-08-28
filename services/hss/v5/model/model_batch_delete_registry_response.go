package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteRegistryResponse Response Object
type BatchDeleteRegistryResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteRegistryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteRegistryResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteRegistryResponse", string(data)}, " ")
}
