package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryDesktopByTagReq struct {

	// 默认为0
	Offset *string `json:"offset,omitempty"`

	// 默认为1000
	Limit *string `json:"limit,omitempty"`

	// 如果是filter就是分页查询，如果是count只需按照条件将总条数返回即可
	Action *string `json:"action,omitempty"`

	// 包含任意一个标签,该字段为true时查询所有不带标签的资源
	WithoutAnyTag *bool `json:"without_any_tag,omitempty"`

	// match对象
	Matches *[]Match `json:"matches,omitempty"`

	// 包含的标签对象，只要有一个不包含，就不符合，一个key对应多个value
	Tags *[]Tags `json:"tags,omitempty"`

	// 包含任意标签，若全都不包含，不符合，一个key对应多个value
	TagsAny *[]Tags `json:"tags_any,omitempty"`

	// 不包含标签，只要有一个不包含，就符合了，一个key对应多个value
	NotTags *[]Tags `json:"not_tags,omitempty"`

	// 不包含任意标签，只要包含其中一个，就不符合，一个key对应多个value
	NotTagsAny *[]Tags `json:"not_tags_any,omitempty"`
}

func (o QueryDesktopByTagReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryDesktopByTagReq struct{}"
	}

	return strings.Join([]string{"QueryDesktopByTagReq", string(data)}, " ")
}
