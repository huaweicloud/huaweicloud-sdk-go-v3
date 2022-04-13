package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
