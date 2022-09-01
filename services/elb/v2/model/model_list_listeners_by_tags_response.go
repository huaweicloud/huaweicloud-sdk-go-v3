package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListListenersByTagsResponse struct {

	// 总记录数。当resources为空时，表示名称为matches字段中指定的value的资源个数；resources不为空时，表示和tags字段匹配的资源的个数。
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count"`

	// 根据tag查询出的资源对象。 当请求中的action为filters，返回体中有该字段。 当请求中的action为count时，返回体中无该字段。
	Resources      *[]ResourcesByTag `json:"resources,omitempty" xml:"resources"`
	HttpStatusCode int               `json:"-"`
}

func (o ListListenersByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListListenersByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListListenersByTagsResponse", string(data)}, " ")
}
