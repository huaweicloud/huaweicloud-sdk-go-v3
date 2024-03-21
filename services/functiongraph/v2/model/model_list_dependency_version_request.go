package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDependencyVersionRequest Request Object
type ListDependencyVersionRequest struct {

	// 依赖包的ID。
	DependId string `json:"depend_id"`

	// 上一次查询依赖包的最后记录位置，默认为\"0\"。
	Marker *string `json:"marker,omitempty"`

	// 单次查询最大条数
	Maxitems *string `json:"maxitems,omitempty"`

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`
}

func (o ListDependencyVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDependencyVersionRequest struct{}"
	}

	return strings.Join([]string{"ListDependencyVersionRequest", string(data)}, " ")
}
