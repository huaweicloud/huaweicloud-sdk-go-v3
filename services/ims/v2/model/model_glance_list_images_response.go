package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type GlanceListImagesResponse struct {
	// 查询首页的URL。

	First *string `json:"first,omitempty"`
	// 资源类型。

	Images *[]GlanceShowImageResponseBody `json:"images,omitempty"`
	// 描述镜像列表模式的URL。

	Schema *string `json:"schema,omitempty"`
	// 查询下一页的URL。当查询镜像列表最后一页时，不存在next。

	Next           *string `json:"next,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o GlanceListImagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceListImagesResponse struct{}"
	}

	return strings.Join([]string{"GlanceListImagesResponse", string(data)}, " ")
}
