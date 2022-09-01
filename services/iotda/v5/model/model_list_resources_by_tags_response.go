package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListResourcesByTagsResponse struct {

	// 资源列表。
	Resources *[]ResourceDto `json:"resources,omitempty" xml:"resources"`

	Page           *Page `json:"page,omitempty" xml:"page"`
	HttpStatusCode int   `json:"-"`
}

func (o ListResourcesByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcesByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListResourcesByTagsResponse", string(data)}, " ")
}
