package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDictionaryResponse struct {

	// 总数，与分页无关
	Total *int32 `json:"total,omitempty" xml:"total"`

	// 当前页的数量，小于等于请求里指定的limit
	Size *int32 `json:"size,omitempty" xml:"size"`

	// 字典列表
	Dictionaries   *[]Dictionary `json:"dictionaries,omitempty" xml:"dictionaries"`
	HttpStatusCode int           `json:"-"`
}

func (o ListDictionaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDictionaryResponse struct{}"
	}

	return strings.Join([]string{"ListDictionaryResponse", string(data)}, " ")
}
