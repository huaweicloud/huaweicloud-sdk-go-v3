package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPiiCategoriesResponse Response Object
type ListPiiCategoriesResponse struct {

	// 分类列表
	Items *[]CategoryItem `json:"items,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPiiCategoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPiiCategoriesResponse struct{}"
	}

	return strings.Join([]string{"ListPiiCategoriesResponse", string(data)}, " ")
}
