package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCategoriesTreeResponse Response Object
type ListCategoriesTreeResponse struct {

	// 目录树信息。
	Categories     *[]CategoryInfo `json:"categories,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListCategoriesTreeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCategoriesTreeResponse struct{}"
	}

	return strings.Join([]string{"ListCategoriesTreeResponse", string(data)}, " ")
}
