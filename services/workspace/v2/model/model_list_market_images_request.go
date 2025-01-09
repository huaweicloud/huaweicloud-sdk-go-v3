package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMarketImagesRequest Request Object
type ListMarketImagesRequest struct {

	// 镜像id，支持传1到100个。
	ImageIds []string `json:"image_ids"`
}

func (o ListMarketImagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMarketImagesRequest struct{}"
	}

	return strings.Join([]string{"ListMarketImagesRequest", string(data)}, " ")
}
