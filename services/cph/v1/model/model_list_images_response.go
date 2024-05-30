package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImagesResponse Response Object
type ListImagesResponse struct {

	// 镜像详情
	ImageInfos *[]ListImagesView `json:"image_infos,omitempty"`

	// 总条数
	Total *int32 `json:"total,omitempty"`

	// 请求的唯一标识ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListImagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImagesResponse struct{}"
	}

	return strings.Join([]string{"ListImagesResponse", string(data)}, " ")
}
