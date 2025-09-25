package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAssetResourceResponse Response Object
type DeleteAssetResourceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAssetResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAssetResourceResponse struct{}"
	}

	return strings.Join([]string{"DeleteAssetResourceResponse", string(data)}, " ")
}
