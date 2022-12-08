package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListImagesResponse struct {

	// 云桌面支持的产品镜像列表。
	Images *[]ImageInfo `json:"images,omitempty"`

	// 对象总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListImagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImagesResponse struct{}"
	}

	return strings.Join([]string{"ListImagesResponse", string(data)}, " ")
}
