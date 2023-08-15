package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishAssetsResponse Response Object
type PublishAssetsResponse struct {

	// 发布的媒资信息。
	AssetInfoArray *[]AssetInfo `json:"asset_info_array,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o PublishAssetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishAssetsResponse struct{}"
	}

	return strings.Join([]string{"PublishAssetsResponse", string(data)}, " ")
}
