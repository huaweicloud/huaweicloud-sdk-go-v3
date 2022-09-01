package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAssociatedWikisResponse struct {

	// 关联的wiki列表
	Wikis *[]AttachWikiDetail `json:"wikis,omitempty" xml:"wikis"`

	// 总数
	Total          *int32 `json:"total,omitempty" xml:"total"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAssociatedWikisResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssociatedWikisResponse struct{}"
	}

	return strings.Join([]string{"ListAssociatedWikisResponse", string(data)}, " ")
}
