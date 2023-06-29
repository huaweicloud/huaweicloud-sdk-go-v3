package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAttributesResponse Response Object
type ListAttributesResponse struct {

	// 每页记录数
	Limit *int64 `json:"limit,omitempty"`

	// 页码
	Offset *int64 `json:"offset,omitempty"`

	// 记录总数
	Count *int64 `json:"count,omitempty"`

	// 自定义属性记录
	Attributes     *[]CmAttributeVo `json:"attributes,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListAttributesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAttributesResponse struct{}"
	}

	return strings.Join([]string{"ListAttributesResponse", string(data)}, " ")
}
