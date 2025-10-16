package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TmsResourceInstancesRequest struct {

	// 索引位置，默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数，默认为1000
	Limit *int32 `json:"limit,omitempty"`

	// 操作标识
	Action *string `json:"action,omitempty"`

	Tags *[]TmsResourceInstancesTag `json:"tags,omitempty"`

	SysTags *[]TmsResourceInstancesTag `json:"sys_tags,omitempty"`

	Matches *[]TmsResourceTag `json:"matches,omitempty"`

	// 无任何标签的资源筛选标识
	WithoutAnyTag *string `json:"without_any_tag,omitempty"`

	TagsAny *[]TmsResourceInstancesTag `json:"tags_any,omitempty"`

	NotTagsAny *[]TmsResourceInstancesTag `json:"not_tags_any,omitempty"`

	NotTags *[]TmsResourceInstancesTag `json:"not_tags,omitempty"`
}

func (o TmsResourceInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TmsResourceInstancesRequest struct{}"
	}

	return strings.Join([]string{"TmsResourceInstancesRequest", string(data)}, " ")
}
