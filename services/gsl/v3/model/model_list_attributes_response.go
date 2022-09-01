package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAttributesResponse struct {

	// 每页记录数
	Limit *int64 `json:"limit,omitempty" xml:"limit"`

	// 页码
	Offset *int64 `json:"offset,omitempty" xml:"offset"`

	// 记录总数
	Count *int64 `json:"count,omitempty" xml:"count"`

	// 自定义属性记录
	Attributes     *[]CmAttributeVo `json:"attributes,omitempty" xml:"attributes"`
	HttpStatusCode int              `json:"-"`
}

func (o ListAttributesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAttributesResponse struct{}"
	}

	return strings.Join([]string{"ListAttributesResponse", string(data)}, " ")
}
