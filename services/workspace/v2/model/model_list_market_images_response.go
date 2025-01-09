package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMarketImagesResponse Response Object
type ListMarketImagesResponse struct {

	// 镜像信息列表。
	Images         *[]Image `json:"images,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ListMarketImagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMarketImagesResponse struct{}"
	}

	return strings.Join([]string{"ListMarketImagesResponse", string(data)}, " ")
}
