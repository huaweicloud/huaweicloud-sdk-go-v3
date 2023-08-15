package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQueryProjectResourceTagsRequest Request Object
type ListQueryProjectResourceTagsRequest struct {

	// 资源类型，值为：endpoint_service或endpoint。
	ResourceType string `json:"resource_type"`

	// 发送的实体的MIME类型。推荐用户默认使用application/json， 如果API是对象、镜像上传等接口，媒体类型可按照流类型的不同进行确定。
	ContentType *string `json:"Content-Type,omitempty"`
}

func (o ListQueryProjectResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueryProjectResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"ListQueryProjectResourceTagsRequest", string(data)}, " ")
}
