package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishRootAssetResponse Response Object
type PublishRootAssetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o PublishRootAssetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishRootAssetResponse struct{}"
	}

	return strings.Join([]string{"PublishRootAssetResponse", string(data)}, " ")
}
