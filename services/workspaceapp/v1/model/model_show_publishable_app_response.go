package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPublishableAppResponse Response Object
type ShowPublishableAppResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 组下面的镜像ID列表。
	GroupImages *[]string `json:"group_images,omitempty"`

	// 查询到的应用列表。
	Items          *[]PublishableApp `json:"items,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowPublishableAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublishableAppResponse struct{}"
	}

	return strings.Join([]string{"ShowPublishableAppResponse", string(data)}, " ")
}
