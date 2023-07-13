package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudImagesResponse Response Object
type ListCloudImagesResponse struct {

	// 查询返回的镜像列表。
	Images         *[]ImageInfo `json:"images,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListCloudImagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudImagesResponse struct{}"
	}

	return strings.Join([]string{"ListCloudImagesResponse", string(data)}, " ")
}
