package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResourceRelationsRequest Request Object
type ShowResourceRelationsRequest struct {

	// 资源ID
	ResourceId string `json:"resource_id"`

	// 资源ID
	RelatedResourceId *string `json:"related_resource_id,omitempty"`

	// 关联资源类型
	RelatedResourceType *string `json:"related_resource_type,omitempty"`

	// 最大的返回数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页
	Marker *string `json:"marker,omitempty"`
}

func (o ShowResourceRelationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceRelationsRequest struct{}"
	}

	return strings.Join([]string{"ShowResourceRelationsRequest", string(data)}, " ")
}
