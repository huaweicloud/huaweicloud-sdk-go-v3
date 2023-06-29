package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCategoryResponse Response Object
type ListCategoryResponse struct {

	// 自定义项列表
	Body           *[]CategoryDetailVo `json:"body,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListCategoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCategoryResponse struct{}"
	}

	return strings.Join([]string{"ListCategoryResponse", string(data)}, " ")
}
